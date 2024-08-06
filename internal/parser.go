package internal

import "fmt"

type NodeType int

const (
	Node_Var NodeType = iota
	Node_And
	Node_Or
	Node_Not
	Node_Implies
	Node_Iff
)

type ASTNode struct {
	Type  NodeType
	Value string
	Left  *ASTNode
	Right *ASTNode
}

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

// NewParser creates a new instance of the Parser struct.
//
// It takes a pointer to a Lexer struct as a parameter and returns a pointer to a Parser struct.
// The function initializes a new Parser struct with the provided lexer and calls the nextToken method to
// advance the lexer to the next token.
// The function then returns the initialized Parser struct.
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

// nextToken advances the parser to the next token.
//
// It calls the NextToken method of the lexer to get the next token and updates the currentToken field of the parser.
// If there is an error while getting the next token, it prints an error message.
func (p *Parser) nextToken() {
	var err error
	p.currentToken, err = p.lexer.NextToken()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func (p *Parser) Parse() *ASTNode {
	return p.parseExpression()
}

func (p *Parser) parseExpression() *ASTNode {
	left := p.parseTerm()
	for p.currentToken.Type == Token_IMPLIES || p.currentToken.Type == Token_IFF {
		op := p.currentToken
		p.nextToken()
		right := p.parseTerm()
		nodeType := Node_Implies
		if op.Type == Token_IFF {
			nodeType = Node_Iff
		}
		left = &ASTNode{
			Type:  nodeType,
			Value: op.Value,
			Left:  left,
			Right: right,
		}
	}

	return left
}

func (p *Parser) parseTerm() *ASTNode {
	left := p.parseFactor()

	for p.currentToken.Type == Token_AND || p.currentToken.Type == Token_OR {
		op := p.currentToken
		p.nextToken()
		right := p.parseFactor()
		nodeType := Node_And
		if op.Type == Token_OR {
			nodeType = Node_Or
		}
		left = &ASTNode{
			Type:  nodeType,
			Value: op.Value,
			Left:  left,
			Right: right,
		}
	}

	return left
}

func (p *Parser) parseFactor() *ASTNode {
	switch p.currentToken.Type {
	case Token_VAR:
		node := &ASTNode{
			Type:  Node_Var,
			Value: p.currentToken.Value,
		}
		p.nextToken()
		return node
	case Token_NOT:
		p.nextToken()
		return &ASTNode{
			Type:  Node_Not,
			Right: p.parseFactor(),
		}
	case Token_LEFT_PAREN:
		p.nextToken()
		node := p.parseExpression()
		if p.currentToken.Type != Token_RIGHT_PAREN {
			panic("Expected closing parenthesis")
		}
		p.nextToken()
		return node
	default:
		panic(fmt.Sprintf("Unexpected token: %v", p.currentToken))
	}
}
