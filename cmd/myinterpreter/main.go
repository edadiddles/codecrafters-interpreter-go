package main

import (
	"fmt"
	"os"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" && command != "parse" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	
    filename := os.Args[2]
    fileContents, err := os.ReadFile(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
        os.Exit(1)
    }
    
    tokens, is_lexical_error := scanner.Scan(fileContents)

    if command == "tokenize" {
        for _, token := range tokens {
            token.PrintToken()
        }
    }

    if command == "parse" {
        parser.Parse(tokens)
    }
    
    status_code := 0
    if is_lexical_error {
        status_code = 65
    }

    os.Exit(status_code)
}
