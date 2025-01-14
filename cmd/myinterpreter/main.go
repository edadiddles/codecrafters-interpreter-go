package main

import (
	"fmt"
	"os"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/evaluator"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" && command != "parse" && command != "evaluate" {
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
        
        if is_lexical_error {
            os.Exit(65)
        } 
    }

    exprs, is_parse_error := parser.Parse(tokens)
    if command == "parse" {
        for _, expr := range exprs {
            str := parser.PrintExpression(expr)
            fmt.Fprintf(os.Stdout, "%s\n", str)
        }

        if is_parse_error {
            os.Exit(65)
        }
    }

    evals, is_runtime_error := evaluator.Evaluate(exprs)
    if command == "evaluate" {
        for _, eval := range evals {
            str := evaluator.PrintEval(eval)
            fmt.Fprintf(os.Stdout, "%s\n", str)
        }

        if is_runtime_error {
            os.Exit(70)
        }
    }

    os.Exit(0)
}
