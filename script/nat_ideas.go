package script

import (
	"fmt"
	"likeawizard/nation-explorer/config"
	"log"
	"os"

	lex "github.com/timtadh/lexmachine"
)

var natIdeaList map[string]NatIdeas

type modType int

const (
	MOD_NUM = iota
	MOD_ID
	MOD_STRING
)

type NatIdeas struct {
	Name       string
	Traditions ModGroup
	Ambition   ModGroup
	Ideas      []ModGroup
}

type Modifier struct {
	Name  string
	Value string
	Type  modType
}

type ModGroup struct {
	Name string
	Mods []Modifier
}

type scriptTree struct {
	parent *scriptTree
	vals   map[string]interface{}
}

func LoadNationalIdeas() {
	natIdeaList = make(map[string]NatIdeas)
	bytes, _ := os.ReadFile(config.NationalIdeaPath)
	isKey := true
	var key string

	curr := &scriptTree{
		parent: nil,
		vals:   make(map[string]interface{}, 0),
	}
	root := curr

	s, err := Lexer.Scanner(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lex.Token)
		if Tokens[token.Type] == "COMMENT" {
			continue
		}

		switch Tokens[token.Type] {
		case "COMMENT":
			continue
		case "}":
			isKey = true
			curr = curr.parent

		case "{":
			isKey = true
			child := &scriptTree{
				parent: curr,
				vals:   make(map[string]interface{}, 0), //TODO
			}
			curr.vals[key] = child
			curr = child
		case "=":
			isKey = false
		case "ID":
			if isKey {
				key = string(token.Lexeme)
			} else {
				curr.vals[key] = string(token.Lexeme)
				isKey = true
			}
		}

	}

	formatIdeas(root)
}

func formatIdeas(root *scriptTree) {
	for country, ideaGroup := range root.vals {
		natIdeas := NatIdeas{Name: country}
		for key, val := range ideaGroup.(*scriptTree).vals {
			switch key {
			case "free", "trigger":
				continue
			case "start":
				natIdeas.Traditions = unpackModifierGroup("Traditions", val.(*scriptTree))
			case "bonus":
				natIdeas.Ambition = unpackModifierGroup("Ambition", val.(*scriptTree))
			default:
				natIdeas.Ideas = append(natIdeas.Ideas, unpackModifierGroup(key, val.(*scriptTree)))
			}
		}
		natIdeaList[country] = natIdeas
	}
}

func GetNatIdeas() map[string]NatIdeas {
	return natIdeaList
}

func unpackModifierGroup(name string, modGroup *scriptTree) ModGroup {
	mg := ModGroup{
		Name: name,
		Mods: make([]Modifier, 0),
	}
	for key, val := range modGroup.vals {
		switch v := val.(type) {
		case string:
			mg.Mods = append(mg.Mods, Modifier{Name: key, Value: v})
		default:
			panic(1)
		}
	}

	return mg
}
