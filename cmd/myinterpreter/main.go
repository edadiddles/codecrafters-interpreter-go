package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	
	 filename := os.Args[2]
	 fileContents, err := os.ReadFile(filename)
	 if err != nil {
	 	fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	 	os.Exit(1)
	 }
	
    for i:=0; i < len(fileContents); i++ {
        if (fileContents[i] == 40) {
            fmt.Println("LEFT_PAREN ( null")
        } else if (fileContents[i] == 41) {
            fmt.Println("RIGHT_PAREN ) null")
        } else if (fileContents[i] == 123) {
            fmt.Println("LEFT_BRACE { null")
        } else if (fileContents[i] == 125) {
            fmt.Println("RIGHT_BRACE } null")
        } else if (fileContents[i] == 44) {
            fmt.Println("COMMA , null")
        } else if (fileContents[i] == 46) {
            fmt.Println("DOT . null")
        } else if (fileContents[i] == 42) {
            fmt.Println("STAR * null")
        } else if (fileContents[i] == 43) {
            fmt.Println("PLUS + null")
        } else if (fileContents[i] == 45) {
            fmt.Println("MINUS - null")
        } else if (fileContents[i] == 59) {
            fmt.Println("SEMICOLON ; null")
        }

    }

    // adding EOF token at end of file
	fmt.Println("EOF  null")
}
