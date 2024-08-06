package internal

import "fmt"

type TruthTable struct {
	Variables []string
	Rows      []map[string]bool
	Results   []bool
}

func GenerateTruthTable(root *ASTNode, variables []string) *TruthTable {
	tt := &TruthTable{
		Variables: variables,
	}

	combinations := 1 << len(variables)
	for i := 0; i < combinations; i++ {
		row := make(map[string]bool)
		for j, v := range variables {
			row[v] = (i & (1 << j)) != 0
		}
		tt.Rows = append(tt.Rows, row)
		tt.Results = append(tt.Results, evaluateExpression(root, row))
	}

	return tt
}

func evaluateExpression(node *ASTNode, values map[string]bool) bool {
	switch node.Type {
	case Node_Var:
		return values[node.Value]
	case Node_And:
		return evaluateExpression(node.Left, values) && evaluateExpression(node.Right, values)
	case Node_Or:
		return evaluateExpression(node.Left, values) || evaluateExpression(node.Right, values)
	case Node_Not:
		return !evaluateExpression(node.Right, values)
	case Node_Implies:
		left := evaluateExpression(node.Left, values)
		right := evaluateExpression(node.Right, values)
		return !left || right
	case Node_Iff:
		left := evaluateExpression(node.Left, values)
		right := evaluateExpression(node.Right, values)
		return left == right
	default:
		panic(fmt.Sprintf("unknown node type: %v", node.Type))
	}
}
