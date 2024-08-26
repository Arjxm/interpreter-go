package main

import (
	"fmt"
	"os"
)

const (
	EOF                   = "EOF"
	LEFT_PAREN            = "LEFT_PAREN"
	RIGHT_PAREN           = "RIGHT_PAREN"
	SEMICOLON             = "SEMICOLON"
	PLUS                  = "PLUS"
	MINUS                 = "MINUS"
	DIVIDE                = "DIVIDE"
	STAR                  = "STAR"
	EQUALS                = "EQUALS"
	ASSIGN                = "ASSIGN"
	NOT_EQUALS            = "NOT_EQUALS"
	LESS_THAN             = "LESS_THAN"
	GREATER_THAN          = "GREATER_THAN"
	LESS_THAN_OR_EQUAL    = "LESS_THAN_OR_EQUAL"
	GREATER_THAN_OR_EQUAL = "GREATER_THAN_OR_EQUAL"
	LEFT_BRACE            = "LEFT_BRACE"
	RIGHT_BRACE           = "RIGHT_BRACE"
	DOT                   = "DOT"
	COMMA                 = "COMMA"
)

var tokenMap = map[byte]string{
	'(': LEFT_PAREN,
	')': RIGHT_PAREN,
	'{': LEFT_BRACE,
	'}': RIGHT_BRACE,
	'+': PLUS,
	'-': MINUS,
	'*': STAR,
	'.': DOT,
	',': COMMA,
	';': SEMICOLON,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	tokenize(filename)
}

func tokenize(filename string) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) == 0 {
		fmt.Printf("%s  null\n", EOF)
		return
	}

	for _, c := range fileContents {
		if tokenType, ok := tokenMap[c]; ok {
			fmt.Printf("%s %s null\n", tokenType, string(c))
		}
	}

	fmt.Printf("%s  null\n", EOF)
}
