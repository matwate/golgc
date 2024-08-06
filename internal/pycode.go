package internal

import "fmt"

func GeneratePythonCode(root *ASTNode) string {
	return generatePythonCode(root)
}

func generatePythonCode(root *ASTNode) string {
	switch root.Type {
	case Node_Var:
		return fmt.Sprintf("Letra('%s')", root.Value)
	case Node_And:
		return fmt.Sprintf("Binario('Y', %s, %s)", generatePythonCode(root.Left), generatePythonCode(root.Right))
	case Node_Or:
		return fmt.Sprintf("Binario('O', %s, %s)", generatePythonCode(root.Left), generatePythonCode(root.Right))
	case Node_Implies:
		return fmt.Sprintf("Binario('>', %s, %s)", generatePythonCode(root.Left), generatePythonCode(root.Right))
	case Node_Not:
		return fmt.Sprintf("Negacion(%s)", generatePythonCode(root.Right))
	case Node_Iff:
		return fmt.Sprintf("Binario('=', %s, %s)", generatePythonCode(root.Left), generatePythonCode(root.Right))
	}
	return ""
}

/**/

