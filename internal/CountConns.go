package internal

func CountConns(root *ASTNode) int {
	return ConnAmount(root)
}

func ConnAmount(node *ASTNode) int {
	if node == nil {
		return 0
	}
	if node.Type == Node_Var {
		return 0
	}
	return 1 + ConnAmount(node.Left) + ConnAmount(node.Right)
}
