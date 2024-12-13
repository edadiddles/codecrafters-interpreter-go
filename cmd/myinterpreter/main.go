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
        // left parenthesis
        if (fileContents[i] == 40) {
            fmt.Println("LEFT_PAREN ( null")
        // right parenthesis
        } else if (fileContents[i] == 41) {
            fmt.Println("RIGHT_PAREN ) null")
        }
    }

    // adding EOF token at end of file
	fmt.Println("EOF  null")
}
