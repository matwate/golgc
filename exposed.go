package main

import "github.com/matwate/lgc/internal"

func Compile(input string) (string, int, string, string) {
	TruthTable, NumberOfConnections, PythonCode, AbstractSyntaxTree := internal.CompileString(input, false)
	return TruthTable, NumberOfConnections, PythonCode, AbstractSyntaxTree
}
