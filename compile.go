package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

// NewLexer creates a new instance of the Lexer struct.
//
// It takes a string input as a parameter and returns a pointer to a Lexer struct.
func NewLexer(input string) *Lexer {
	return &Lexer{bufio.NewReader(strings.NewReader(input))}
}

// NextToken returns the next token from the input stream.
//
// It reads runes from the lexer's reader and returns a Token struct and an error.
// If the end of the input stream is reached, it returns a Token with TokenType Token_EOF and an empty string.
// If an error occurs while reading the input stream, it returns an empty Token and the error.
// If the rune is a space, it continues to the next rune.
// If the rune is a letter or a digit, it unreads the rune and returns the result of calling lexVariable().
// If the rune is '*', it returns a Token with TokenType Token_AND and the string representation of the rune.
// If the rune is '+', it returns a Token with TokenType Token_OR and the string representation of the rune.
// If the rune is '!', it returns a Token with TokenType Token_NOT and the string representation of the rune.
// If the rune is '=', it checks if the next rune is '>'. If it is, it reads the next rune and returns a Token with TokenType Token_IMPLIES and the string "=>".
// If the rune is '(', it returns a Token with TokenType Token_LEFT_PAREN and the string representation of the rune.
// If the rune is ')', it returns a Token with TokenType Token_RIGHT_PAREN and the string representation of the rune.
// If the rune is '<', it checks if the next two runes are '=' and '>'. If they are, it reads the next rune and returns a Token with TokenType Token_IFF and the string "<=>".
// If none of the above conditions are met, it returns an empty Token and an error indicating an unexpected character.
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
		case unicode.IsLetter(r) || unicode.IsDigit(r):
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

// lexVariable scans the input for a variable name and returns a Token representing the variable.
//
// It reads runes from the input until it encounters a non-letter or non-digit character. It then
// unreads the last character and returns a Token with the type Token_VAR and the scanned variable
// name as the value. If an error occurs while reading the input, it returns an empty Token and the
// error.
//
// Parameters:
// - l: A pointer to a Lexer struct representing the lexer.
//
// Returns:
// - Token: A Token representing the variable.
// - error: An error if an error occurred while reading the input.
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
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
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

func CompileString(input string, trueOnly bool) {

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	ast := parser.Parse()

	//fmt.Println("Original AST:")
	//printAST(ast, 0)

	analyzer := NewSemanticAnalyzer()
	if err := analyzer.Analyze(ast); err != nil {
		//fmt.Printf("Semantic error: %v\n", err)
		return
	}

	variables := analyzer.GetVariables()
	//fmt.Printf("Variables: %v\n", variables)

	simplifiedAST := Simplify(ast)
	//fmt.Println("\nSimplified AST:")
	//printAST(simplifiedAST, 0)

	tt := GenerateTruthTable(simplifiedAST, variables)
	fmt.Println("\nTruth Table:")

	conns := CountConns(simplifiedAST)

	if !trueOnly {
		printTruthTable(tt)
	} else {
		printTruthTableTrueOnly(tt)
	}

	printConns(conns)
}

func printTruthTable(tt *TruthTable) {
	file, err := os.Create(".lgout")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range tt.Variables {
		fmt.Printf("%s\t", v)
		fmt.Fprintf(writer, "%s\t", v)
	}
	fmt.Println("Result")
	fmt.Fprintf(writer, "Result\n")

	for i, row := range tt.Rows {
		for _, v := range tt.Variables {
			if row[v] {
				fmt.Print("T\t")
				fmt.Fprintf(writer, "T\t")
			} else {
				fmt.Print("F\t")
				fmt.Fprintf(writer, "F\t")
			}
		}
		if tt.Results[i] {
			fmt.Println("T")
			fmt.Fprintf(writer, "T\n")
		} else {
			fmt.Println("F")
			fmt.Fprintf(writer, "F\n")
		}
	}

	writer.Flush()

}

func printTruthTableTrueOnly(tt *TruthTable) {
	file, err := os.Create(".lgout")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range tt.Variables {
		fmt.Printf("%s\t", v)
		fmt.Fprintf(writer, "%s\t", v)
	}
	fmt.Println("Result")
	fmt.Fprintf(writer, "Result\n")

	for i, row := range tt.Rows {
		if !tt.Results[i] {
			continue
		}
		for _, v := range tt.Variables {
			if row[v] {
				fmt.Print("T\t")
				fmt.Fprintf(writer, "T\t")
			} else {
				fmt.Print("F\t")
				fmt.Fprintf(writer, "F\t")
			}
		}
		if tt.Results[i] {
			fmt.Println("T")
			fmt.Fprintf(writer, "T\n")
		} else {
			fmt.Println("F")
			fmt.Fprintf(writer, "F\n")
		}
	}

	writer.Flush()

}

func printAST(node *ASTNode, indent int) {
	if node == nil {
		return
	}
	fmt.Printf("%s%v: %s\n", strings.Repeat("  ", indent), node.Type, node.Value)
	printAST(node.Left, indent+1)
	printAST(node.Right, indent+1)
}

func printConns(conns int) {

	file, err := os.Open(".lgout")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Println("Connections: ", conns)
	fmt.Fprintf(writer, "Connections: %d\n", conns)

	writer.Flush()
}
