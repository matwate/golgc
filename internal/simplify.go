package internal

func Simplify(node *ASTNode) *ASTNode {
	if node == nil {
		return nil
	}

	switch node.Type {
	case Node_And, Node_Or, Node_Implies, Node_Iff:
		node.Left = Simplify(node.Left)
		node.Right = Simplify(node.Right)
		return simplifyBinaryOp(node)
	case Node_Not:
		node.Right = Simplify(node.Right)
		return simplifyNot(node)
	default:
		return node
	}
}

func simplifyBinaryOp(node *ASTNode) *ASTNode {
	// Implement simplification rules for binary operations
	// For example: a + a = a, a | a = a, etc.
	if node.Left.Type == Node_Var && node.Right.Type == Node_Var && node.Left.Value == node.Right.Value {
		if node.Type == Node_And || node.Type == Node_Or {
			return node.Left
		}

		if node.Type == Node_Implies {
			return node.Left
		}

		if node.Type == Node_Iff {
			return &ASTNode{Type: Node_Var, Value: "true"}
		}
	}
	return node
}

func simplifyNot(node *ASTNode) *ASTNode {
	// Implement simplification rules for NOT
	// For example: !!a = a
	if node.Right.Type == Node_Not {
		return node.Right.Right
	}
	return node
}
