package golgc

import "github.com/matwate/golgc/internal"

func Compile(input string) (string, int, string, string) {
	TruthTable, NumberOfConnections, PythonCode, AbstractSyntaxTree := internal.CompileString(input, false)
	return TruthTable, NumberOfConnections, PythonCode, AbstractSyntaxTree
}
