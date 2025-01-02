package main

import (
	"fmt"
	"os"
    "strconv"
)


//type Token struct {
//    type []u8,
//    lexeme []u8,
//    literal []u8,
//}


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

    line_num := 1
    is_lexical_error := false
    for i:=0; i < len(fileContents); i++ {
        if (fileContents[i] == '(') {
            fmt.Println("LEFT_PAREN ( null")
        } else if (fileContents[i] == ')') {
            fmt.Println("RIGHT_PAREN ) null")
        } else if (fileContents[i] == '{') {
            fmt.Println("LEFT_BRACE { null")
        } else if (fileContents[i] == '}') {
            fmt.Println("RIGHT_BRACE } null")
        } else if (fileContents[i] == ',') {
            fmt.Println("COMMA , null")
        } else if (fileContents[i] == '.') {
            fmt.Println("DOT . null")
        } else if (fileContents[i] == '*') {
            fmt.Println("STAR * null")
        } else if (fileContents[i] == '+') {
            fmt.Println("PLUS + null")
        } else if (fileContents[i] == '-') {
            fmt.Println("MINUS - null")
        } else if (fileContents[i] == ';') {
            fmt.Println("SEMICOLON ; null")
        } else if fileContents[i] == '!' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                fmt.Println("BANG_EQUAL != null")
                i++
            } else {
                fmt.Println("BANG ! null")
            }
        } else if fileContents[i] == '=' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                fmt.Println("EQUAL_EQUAL == null")
                i++
            } else {
                fmt.Println("EQUAL = null")
            }
        } else if fileContents[i] == '<' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                fmt.Println("LESS_EQUAL <= null")
                i++
            } else {
                fmt.Println("LESS < null")
            }
        } else if fileContents[i] == '>' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                fmt.Println("GREATER_EQUAL >= null")
                i++
            } else {
                fmt.Println("GREATER > null")
            }
        } else if fileContents[i] == '/' {
            if len(fileContents) > i+1 && fileContents[i+1] == '/' {
                i++
                for ; i < len(fileContents) && fileContents[i] != '\n'; i++ {} // do nothing
                line_num++
            } else {
                fmt.Println("SLASH / null")
            }
        } else if fileContents[i] == '"' {
            i++
            contents := []byte{}
            for ; i < len(fileContents) && fileContents[i] != '"'; i++ {
                contents = append(contents, fileContents[i])
            }
            if i < len(fileContents) && fileContents[i] == '"' {
                fmt.Printf("STRING \"%s\" %s\n", string(contents), string(contents))
            } else {
                is_lexical_error = true
                fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n",  line_num)
            }
        } else if fileContents[i] >= '0' && fileContents[i] <= '9' {
            contents := []byte{}
            has_decimal := false
            num_decimal_digits := 0
            has_non_zero_decimal := false
            for ; i < len(fileContents) && ((fileContents[i] >= '0' && fileContents[i] <= '9') || fileContents[i] == '.'); i++ {
                if fileContents[i] == '.' {
                    if has_decimal {
                        fmt.Fprintf(os.Stderr, "[line %d] Error: ...\n", line_num)
                    }
                    
                    has_decimal = true
                    if len(fileContents) > 0 {
                        if len(fileContents) > i+1 && fileContents[i+1] >= '0' && fileContents[i+1] <= '9' {
                            contents = append(contents, fileContents[i])
                            i++
                            contents = append(contents, fileContents[i])
                            if fileContents[i] != '0'{
                                has_non_zero_decimal = true
                            }
                            num_decimal_digits++
                        } else {
                            fmt.Fprintf(os.Stderr, "[line %d] Error: ...\n", line_num)
                        }
                    } else {
                        fmt.Fprintf(os.Stderr, "[line %d] Error: ...\n", line_num)
                    }
                } else {
                    contents = append(contents, fileContents[i])
                    if has_decimal { 
                        if fileContents[i] != '0'{
                            has_non_zero_decimal = true
                        }
                        num_decimal_digits++
                    }
                }
            }
            if num_decimal_digits == 0 || !has_non_zero_decimal {
                num_decimal_digits = 1
            }

            f, _ := strconv.ParseFloat(string(contents), 32)
            float_fmt := fmt.Sprintf("%s%df", "%.", num_decimal_digits)
            fmt_str := fmt.Sprintf("NUMBER %s %s\n", string(contents), float_fmt)
            fmt.Printf(fmt_str, f)
            i-- //decrement counter to avoid skipping due to increment in parent for loop
        } else if (fileContents[i] >= 'a' && fileContents[i] <= 'z') || (fileContents[i] >= 'A' && fileContents[i] <= 'Z') || (fileContents[i] == '_') || (fileContents[i] >= '0' && fileContents[i] <= '9') {
            word := []byte{}
            word = append(word, fileContents[i])
            for i++; i < len(fileContents) && ((fileContents[i] >= 'a' && fileContents[i] <= 'z') || (fileContents[i] >= 'A' && fileContents[i] <= 'Z') || (fileContents[i] == '_') || (fileContents[i] >= '0' && fileContents[i] <= '9')); i++ {
                word = append(word, fileContents[i])
            }
            i--

            if string(word) == "and" {
                fmt.Println("AND and null")
            } else if string(word) == "class" {
                fmt.Println("CLASS class null")
            } else if string(word) == "else" {
                fmt.Println("ELSE else null")
            } else if string(word) == "false" {
                fmt.Println("FALSE false null")
            } else if string(word) == "for" {
                fmt.Println("FOR for null")
            } else if string(word) == "fun" {
                fmt.Println("FUN fun null")
            } else if string(word) == "if" {
                fmt.Println("IF if null")
            } else if string(word) == "nil" {
                fmt.Println("NIL nil null")
            } else if string(word) == "or" {
                fmt.Println("OR or null")
            } else if string(word) == "print" {
                fmt.Println("PRINT print null")
            } else if string(word) == "return" {
                fmt.Println("RETURN return null")
            } else if string(word) == "super" {
                fmt.Println("SUPER super null")
            } else if string(word) == "this" {
                fmt.Println("THIS this null")
            } else if string(word) == "true" {
                fmt.Println("TRUE true null")
            } else if string(word) == "var" {
                fmt.Println("VAR var null")
            } else if string(word) == "while" {
                fmt.Println("WHILE while null")
            } else {
                fmt.Printf("IDENTIFIER %s null\n", word)
            }
         
        } else if fileContents[i] == ' ' {
            // do nothing
        } else if fileContents[i] == '\t' {
            // do nothing
        } else if fileContents[i] == '\n' {
            line_num++
        } else if (fileContents[i] == '$') {
            is_lexical_error = true
            fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: $\n",  line_num)
        } else if (fileContents[i] == '#') {
            is_lexical_error = true
            fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: #\n", line_num)
        } else {
            is_lexical_error = true
            fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line_num, fileContents[i])
        }
    }

    // adding EOF token at end of file
	fmt.Println("EOF  null")

    status_code := 0
    if is_lexical_error {
        status_code = 65
    }

    os.Exit(status_code)
}
