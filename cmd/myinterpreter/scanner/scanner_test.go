package scanner

import (
    "fmt"
    "testing"
)


func TestBasicScanner(t *testing.T) {
    input := []byte{ '{', '(', ',', '*', '.', ')', '}' }

    tokens, _ := Scan(input)

    fmt.Printf("%v", tokens)
}
