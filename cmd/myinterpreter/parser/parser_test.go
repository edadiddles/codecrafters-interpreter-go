package parser

import (
    "testing"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

func TestBasicParser(t *testing.T) {
    input := []byte{ 't', 'r', 'u', 'e' }
    tokens, _ := scanner.Scan(input)

    Parse(tokens)
}
