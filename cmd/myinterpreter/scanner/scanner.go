package scanner

import (
    "strconv"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)


func Scan(fileContents []byte) {
    line_num := 1
    is_lexical_error := false
    for i:=0; i < len(fileContents); i++ {
        if (fileContents[i] == '(') {
            t := token.Token[string]{ Type: token.LeftParen, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == ')') {
            t := token.Token[string]{ Type: token.RightParen, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '{') {
            t := token.Token[string]{ Type: token.LeftBrace, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '}') {
            t := token.Token[string]{ Type: token.RightBrace, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == ',') {
            t := token.Token[string]{ Type: token.Comma, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '.') {
            t := token.Token[string]{ Type: token.Period, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '*') {
            t := token.Token[string]{ Type: token.Asterisk, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '+') {
            t := token.Token[string]{ Type: token.Plus, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '-') {
            t := token.Token[string]{ Type: token.Minus, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == ';') {
            t := token.Token[string]{ Type: token.Semicolon, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
        } else if fileContents[i] == '!' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t := token.Token[string]{ Type: token.BangEqual, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
                i++
            } else {
                t := token.Token[string]{ Type: token.Bang, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
            }
        } else if fileContents[i] == '=' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t := token.Token[string]{ Type: token.EqualEqual, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
                i++
            } else {
                t := token.Token[string]{ Type: token.Equal, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
            }
        } else if fileContents[i] == '<' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t := token.Token[string]{ Type: token.LessEqual, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
                i++
            } else {
                t := token.Token[string]{ Type: token.Less, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
            }
        } else if fileContents[i] == '>' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t := token.Token[string]{ Type: token.GreaterEqual, Lexeme: string(fileContents[i:i+2]), Literal: "null", Line: line_num }
            t.PrintToken()
                i++
            } else {
                t := token.Token[string]{ Type: token.Greater, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
            }
        } else if fileContents[i] == '/' {
            if len(fileContents) > i+1 && fileContents[i+1] == '/' {
                i++
                for ; i < len(fileContents) && fileContents[i] != '\n'; i++ {} // do nothing
                line_num++
            } else {
                t := token.Token[string]{ Type: token.Slash, Lexeme: string(fileContents[i]), Literal: "null", Line: line_num }
            t.PrintToken()
            }
        } else if fileContents[i] == '"' {
            i++
            contents := []byte{}
            for ; i < len(fileContents) && fileContents[i] != '"'; i++ {
                contents = append(contents, fileContents[i])
            }
            if i < len(fileContents) && fileContents[i] == '"' {
                t := token.Token[string]{ Type: token.String, Lexeme: string(contents), Literal: string(contents), Line: line_num }
            t.PrintToken()
            } else {
                is_lexical_error = true
                t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "Unterminated String", Line: line_num }
            t.PrintToken()
            }
        } else if IsNumeric(fileContents[i]) {
            contents := []byte{}
            has_decimal := false
            num_decimal_digits := 0
            has_non_zero_decimal := false
            for ; i < len(fileContents) && (IsNumeric(fileContents[i]) || fileContents[i] == '.'); i++ {
                if fileContents[i] == '.' {
                    if has_decimal {
                        t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "...", Line: line_num }
            t.PrintToken()
                    }
                    
                    has_decimal = true
                    if len(fileContents) > 0 {
                        if len(fileContents) > i+1 && IsNumeric(fileContents[i]) {
                            contents = append(contents, fileContents[i])
                            i++
                            contents = append(contents, fileContents[i])
                            if fileContents[i] != '0'{
                                has_non_zero_decimal = true
                            }
                            num_decimal_digits++
                        } else {
                            t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "...", Line: line_num }
            t.PrintToken()
                        }
                    } else {
                        t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "...", Line: line_num }
            t.PrintToken()
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

            f, _ := strconv.ParseFloat(string(contents), 64)
            t := token.Token[float64]{ Type: token.Number, Lexeme: string(contents), Literal: f, Line: line_num }
            t.PrintToken()
            i-- //decrement counter to avoid skipping due to increment in parent for loop
        } else if IsAlpha(fileContents[i]) || IsNumeric(fileContents[i]) {
            word := []byte{}
            word = append(word, fileContents[i])
            for i=i+1; i < len(fileContents) && (IsAlpha(fileContents[i]) || IsNumeric(fileContents[i])); i++ {
                word = append(word, fileContents[i])
            }
            i--

            if string(word) == "and" {
                t := token.Token[string]{ Type: token.And, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "class" {
                t := token.Token[string]{ Type: token.Class, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "else" {
                t := token.Token[string]{ Type: token.Else, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "false" {
                t := token.Token[string]{ Type: token.False, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "for" {
                t := token.Token[string]{ Type: token.For, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "fun" {
                t := token.Token[string]{ Type: token.Fun, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "if" {
                t := token.Token[string]{ Type: token.If, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "nil" {
                t := token.Token[string]{ Type: token.Nil, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "or" {
                t := token.Token[string]{ Type: token.Or, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "print" {
                t := token.Token[string]{ Type: token.Print, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "return" {
                t := token.Token[string]{ Type: token.Return, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "super" {
                t := token.Token[string]{ Type: token.Super, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "this" {
                t := token.Token[string]{ Type: token.This, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "true" {
                t := token.Token[string]{ Type: token.True, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "var" {
                t := token.Token[string]{ Type: token.Var, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else if string(word) == "while" {
                t := token.Token[string]{ Type: token.While, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            } else {
                t := token.Token[string]{ Type: token.Identifier, Lexeme: string(word), Literal: "null", Line: line_num }
            t.PrintToken()
            }
         
        } else if fileContents[i] == ' ' {
            // do nothing
        } else if fileContents[i] == '\t' {
            // do nothing
        } else if fileContents[i] == '\n' {
            line_num++
        } else if (fileContents[i] == '$') {
            is_lexical_error = true
            t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "Unexpected character", Line: line_num }
            t.PrintToken()
        } else if (fileContents[i] == '#') {
            is_lexical_error = true
            t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "Unexpected character", Line: line_num }
            t.PrintToken()
        } else {
            is_lexical_error = true
            t := token.Token[string]{ Type: token.Error, Lexeme: string(fileContents[i]), Literal: "Unexpected character", Line: line_num }
            t.PrintToken()
        }
    }

    // adding EOF token at end of file
    t := token.Token[string]{ Type: token.EOF, Lexeme: string(""), Literal: "null", Line: line_num }
    t.PrintToken()
}

func Peek() {}

func Advanced() {}

func IsNumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_') 
}

func IsAlpha(c byte) bool {
    return c >= '0' && c <= '9' 
}
