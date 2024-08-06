package internal

import (
	"fmt"
	"sort"
)

type SemanticAnalyzer struct {
	variables map[string]bool
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{
		variables: make(map[string]bool),
	}
}

func (sa *SemanticAnalyzer) Analyze(node *ASTNode) error {
	switch node.Type {
	case Node_Var:
		sa.variables[node.Value] = true
	case Node_And, Node_Or, Node_Implies:
		if err := sa.Analyze(node.Left); err != nil {
			return err
		}
		if err := sa.Analyze(node.Right); err != nil {
			return err
		}
	case Node_Not:
		if err := sa.Analyze(node.Right); err != nil {
			return err
		}
	case Node_Iff:
		if err := sa.Analyze(node.Left); err != nil {
			return err
		}
		if err := sa.Analyze(node.Right); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown node type: %v", node.Type)
	}
	return nil
}

func (sa *SemanticAnalyzer) GetVariables() []string {
	vars := make([]string, 0, len(sa.variables))
	for v := range sa.variables {
		vars = append(vars, v)
	}
	sort.Strings(vars)
	return vars
}
