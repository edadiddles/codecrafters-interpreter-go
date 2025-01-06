package scanner

import (
    "fmt"
    "testing"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)


func TestBasicScanner(t *testing.T) {
    input := []byte{ '{', '(', ',', '*', '.', ')', '}' }

    tokens := Scan(input)

    fmt.Printf("%v", tokens)
}
