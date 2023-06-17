package script

import (
	"fmt"
	"log"

	lex "github.com/timtadh/lexmachine"
)

type modType int

const (
	MOD_NUM = iota
	MOD_ID
	MOD_STRING
)

type NatIdeas struct {
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
	Mods []Modifier
}

type ideaTree struct {
	parent *ideaTree
	vals   map[string]interface{}
}

func ParseIdeas(bytes []byte) map[string]NatIdeas {
	isKey := true
	var key string

	curr := &ideaTree{
		parent: nil,
		vals:   make(map[string]interface{}, 0),
	}
	root := curr

	s, err := Lexer.Scanner(bytes)
	if err != nil {
		fmt.Println(err)
		return nil
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
			child := &ideaTree{
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

	return formatIdeas(root)
}

func formatIdeas(root *ideaTree) map[string]NatIdeas {
	allIdeas := make(map[string]NatIdeas, 0)

	for country, ideaGroup := range root.vals {
		natIdeas := NatIdeas{}
		for key, val := range ideaGroup.(*ideaTree).vals {
			switch key {
			case "free", "trigger":
				continue
			case "start":
				natIdeas.Traditions = unpackModifierGroup(val.(*ideaTree))
			case "bonus":
				natIdeas.Ambition = unpackModifierGroup(val.(*ideaTree))
			default:
				natIdeas.Ideas = append(natIdeas.Ideas, unpackModifierGroup(val.(*ideaTree)))
			}
		}
		allIdeas[country] = natIdeas
	}

	return allIdeas
}

func unpackModifierGroup(modGroup *ideaTree) ModGroup {
	mg := ModGroup{
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
