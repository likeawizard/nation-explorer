package script

import (
	"fmt"
	"log"
	"strings"

	lex "github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var Literals []string       // The tokens representing literal strings
var Keywords []string       // The keyword tokens
var Tokens []string         // All of the tokens (including literals and keywords)
var TokenIds map[string]int // A map from the token names to their int ids
var Lexer *lex.Lexer        // The lexer object. Use this to construct a Scanner

func init() {
	initTokens()
	Lexer = initLexer()
}

func initTokens() {
	Literals = []string{
		"{",
		"}",
		"=",
		`"`,
	}

	Tokens = []string{
		"COMMENT",
		"ID",
	}
	Tokens = append(Tokens, Literals...)
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

func initLexer() *lex.Lexer {
	lexer := lex.NewLexer()
	for _, lit := range Literals {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(lit))
	}
	for _, name := range Keywords {
		lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	lexer.Add([]byte(`#[^\n]*\n?`), token("COMMENT"))
	lexer.Add([]byte(`([a-z]|[A-Z])([a-z]|[A-Z]|[0-9]|_|-)*`), token("ID")) // alpha-numerics with underscore
	lexer.Add([]byte(`-?([0-9])(.?[0-9]+)*`), token("ID"))                  // numerics - positive/negative floats or ints
	lexer.Add([]byte("( |\t|\n|\r)+"), skip)

	err := lexer.Compile()
	if err != nil {
		panic(err)
	}
	return lexer
}

func skip(*lex.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

func token(name string) lex.Action {
	return func(s *lex.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}

func Parse(bytes []byte) {
	s, err := Lexer.Scanner(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Type    | Lexeme     | Position")
	fmt.Println("--------+------------+------------")

	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lex.Token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n",
			Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
}
