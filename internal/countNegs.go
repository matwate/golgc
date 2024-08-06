package internal

func CountNegs(root *ASTNode) int {
	return countNegs(root)
}

func countNegs(root *ASTNode) int {
	switch root.Type {
	case Node_Var:
		return 0
	case Node_And:
		return countNegs(root.Left) + countNegs(root.Right)
	case Node_Or:
		return countNegs(root.Left) + countNegs(root.Right)
	case Node_Implies:
		return countNegs(root.Left) + countNegs(root.Right)
	case Node_Not:
		return 1 + countNegs(root.Right)
	case Node_Iff:
		return countNegs(root.Left) + countNegs(root.Right)
	}
	return 0
}
