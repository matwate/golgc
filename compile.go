package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	Token_VAR TokenType = iota
	Token_AND
	Token_OR
	Token_NOT
	Token_IMPLIES

	Token_LEFT_PAREN
	Token_RIGHT_PAREN
	Token_EOF
	Token_IFF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	reader *bufio.Reader
}

func NewLexer(input string) *Lexer {
	return &Lexer{bufio.NewReader(strings.NewReader(input))}
}

func (l *Lexer) NextToken() (Token, error) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{Token_EOF, ""}, nil
			}
			return Token{}, err
		}

		switch {
		case unicode.IsSpace(r):
			continue
		case unicode.IsLetter(r):
			l.reader.UnreadRune()
			return l.lexVariable()
		case r == '*':
			return Token{Token_AND, string(r)}, nil
		case r == '+':
			return Token{Token_OR, string(r)}, nil
		case r == '!':
			return Token{Token_NOT, string(r)}, nil
		case r == '=':
			if l.peek() == '>' {
				l.reader.ReadRune()
				return Token{Token_IMPLIES, "=>"}, nil
			}
		case r == '(':
			return Token{Token_LEFT_PAREN, string(r)}, nil
		case r == ')':
			return Token{Token_RIGHT_PAREN, string(r)}, nil
		case r == '<':
			if l.peek() == '=' {
				l.reader.ReadRune()
				if l.peek() == '>' {
					l.reader.ReadRune()
					return Token{Token_IFF, "<=>"}, nil
				}
			}
		default:
			return Token{}, fmt.Errorf("unexpected character: %c", r)
		}
	}
}

func (l *Lexer) lexVariable() (Token, error) {
	var sb strings.Builder
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{}, err
		}
		if !unicode.IsLetter(r) {
			l.reader.UnreadRune()
			break
		}
		sb.WriteRune(r)
	}

	return Token{Token_VAR, sb.String()}, nil
}

func (l *Lexer) peek() rune {
	r, _, _ := l.reader.ReadRune()
	l.reader.UnreadRune()
	return r
}

func CompileString(input string) {

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	ast := parser.Parse()

	fmt.Println("Original AST:")
	printAST(ast, 0)

	analyzer := NewSemanticAnalyzer()
	if err := analyzer.Analyze(ast); err != nil {
		fmt.Printf("Semantic error: %v\n", err)
		return
	}

	variables := analyzer.GetVariables()
	fmt.Printf("Variables: %v\n", variables)

	simplifiedAST := Simplify(ast)
	fmt.Println("\nSimplified AST:")
	printAST(simplifiedAST, 0)

	tt := GenerateTruthTable(simplifiedAST, variables)
	fmt.Println("\nTruth Table:")
	printTruthTable(tt)
}

func printTruthTable(tt *TruthTable) {
	for _, v := range tt.Variables {
		fmt.Printf("%s\t", v)
	}
	fmt.Println("Result")

	for i, row := range tt.Rows {
		for _, v := range tt.Variables {
			if row[v] {
				fmt.Print("T\t")
			} else {
				fmt.Print("F\t")
			}
		}
		if tt.Results[i] {
			fmt.Println("T")
		} else {
			fmt.Println("F")
		}
	}
}

func printAST(node *ASTNode, indent int) {
	if node == nil {
		return
	}
	fmt.Printf("%s%v: %s\n", strings.Repeat("  ", indent), node.Type, node.Value)
	printAST(node.Left, indent+1)
	printAST(node.Right, indent+1)
}
