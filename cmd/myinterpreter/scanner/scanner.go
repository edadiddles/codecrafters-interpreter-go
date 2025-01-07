package scanner

import (
    "strconv"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)


func Scan(fileContents []byte) ([]*token.Token, bool) {
    tokens := make([]*token.Token, 0)
    line_num := 1
    is_lexical_error := false
    var t *token.Token
    for i:=0; i < len(fileContents); i++ {
        if (fileContents[i] == '(') {
            t = token.CreateToken(token.LeftParen, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == ')') {
            t = token.CreateToken(token.RightParen, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '{') {
            t = token.CreateToken(token.LeftBrace, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '}') {
            t = token.CreateToken(token.RightBrace, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == ',') {
            t = token.CreateToken(token.Comma, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '.') {
            t = token.CreateToken(token.Period, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '*') {
            t = token.CreateToken(token.Asterisk, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '+') {
            t = token.CreateToken(token.Plus, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '-') {
            t = token.CreateToken(token.Minus, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == ';') {
            t = token.CreateToken(token.Semicolon, string(fileContents[i]), "null", line_num)
            tokens = append(tokens, t)
        } else if fileContents[i] == '!' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t = token.CreateToken(token.BangEqual, string(fileContents[i:i+2]), "null", line_num)
                tokens = append(tokens, t)
                i++
            } else {
                t = token.CreateToken(token.Bang, string(fileContents[i]), "null", line_num)
                tokens = append(tokens, t)
            }
        } else if fileContents[i] == '=' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t = token.CreateToken(token.EqualEqual, string(fileContents[i:i+2]), "null", line_num)
                tokens = append(tokens, t)
                i++
            } else {
                t = token.CreateToken(token.Equal, string(fileContents[i]), "null", line_num)
                tokens = append(tokens, t)
            }
        } else if fileContents[i] == '<' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t = token.CreateToken(token.LessEqual, string(fileContents[i:i+2]), "null", line_num)
                tokens = append(tokens, t)
                i++
            } else {
                t = token.CreateToken(token.Less, string(fileContents[i]), "null", line_num)
                tokens = append(tokens, t)
            }
        } else if fileContents[i] == '>' {
            if len(fileContents) > i+1 && fileContents[i+1] == '='{
                t = token.CreateToken(token.GreaterEqual, string(fileContents[i:i+2]), "null", line_num)
                tokens = append(tokens, t)
                i++
            } else {
                t = token.CreateToken(token.Greater, string(fileContents[i]), "null", line_num)
                tokens = append(tokens, t)
            }
        } else if fileContents[i] == '/' {
            if len(fileContents) > i+1 && fileContents[i+1] == '/' {
                i++
                for ; i < len(fileContents) && fileContents[i] != '\n'; i++ {} // do nothing
                line_num++
            } else {
                t = token.CreateToken(token.Slash, string(fileContents[i]), "null", line_num)
                tokens = append(tokens, t)
            }
        } else if fileContents[i] == '"' {
            i++
            contents := []byte{}
            for ; i < len(fileContents) && fileContents[i] != '"'; i++ {
                contents = append(contents, fileContents[i])
            }
            if i < len(fileContents) && fileContents[i] == '"' {
                t = token.CreateToken(token.String, string(contents), string(contents), line_num)
                tokens = append(tokens, t)
            } else {
                is_lexical_error = true
                t = token.CreateToken(token.UnterminatedStringError, string(contents), "", line_num)
                tokens = append(tokens, t)
            }
        } else if IsNumeric(fileContents[i]) {
            contents := []byte{}
            has_decimal := false
            for ; i < len(fileContents) && (IsNumeric(fileContents[i]) || fileContents[i] == '.'); i++ {
                if fileContents[i] == '.' {
                    // has mulitple decimals in number
                    if has_decimal {
                        t = token.CreateToken(token.NumericError, string(contents), "", line_num)
                        tokens = append(tokens, t)
                        break
                    // Number starts with decimal
                    } else if len(contents) == 0 { 
                        t = token.CreateToken(token.NumericError, string(contents), "", line_num)
                        tokens = append(tokens, t)
                        break
                    }
                    
                    has_decimal = true
                    if len(fileContents) > i+1 && IsNumeric(fileContents[i+1]) {
                        contents = append(contents, fileContents[i])
                        i++
                        contents = append(contents, fileContents[i])
                    } else {
                        t = token.CreateToken(token.NumericError, string(contents), "", line_num)
                        tokens = append(tokens, t)
                    }
                } else {
                    contents = append(contents, fileContents[i])
                }
            }

            f, _ := strconv.ParseFloat(string(contents), 64)
            t = token.CreateToken(token.Number, string(contents), f, line_num)
            tokens = append(tokens, t)
            i-- //decrement counter to avoid skipping due to increment in parent for loop
        } else if IsAlpha(fileContents[i]) || IsNumeric(fileContents[i]) {
            word := []byte{}
            word = append(word, fileContents[i])
            for i=i+1; i < len(fileContents) && (IsAlpha(fileContents[i]) || IsNumeric(fileContents[i])); i++ {
                word = append(word, fileContents[i])
            }
            i--

            if string(word) == "and" {
                t = token.CreateToken(token.And, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "class" {
                t = token.CreateToken(token.Class, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "else" {
                t = token.CreateToken(token.Else, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "false" {
                t = token.CreateToken(token.False, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "for" {
                t = token.CreateToken(token.For, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "fun" {
                t = token.CreateToken(token.Fun, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "if" {
                t = token.CreateToken(token.If, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "nil" {
                t = token.CreateToken(token.Nil, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "or" {
                t = token.CreateToken(token.Or, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "print" {
                t = token.CreateToken(token.Print, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "return" {
                t = token.CreateToken(token.Return, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "super" {
                t = token.CreateToken(token.Super, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "this" {
                t = token.CreateToken(token.This, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "true" {
                t = token.CreateToken(token.True, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "var" {
                t = token.CreateToken(token.Var, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else if string(word) == "while" {
                t = token.CreateToken(token.While, string(word), "null", line_num)
                tokens = append(tokens, t)
            } else {
                t = token.CreateToken(token.Identifier, string(word), "null", line_num)
                tokens = append(tokens, t)
            }
        } else if fileContents[i] == ' ' {
            // do nothing
        } else if fileContents[i] == '\t' {
            // do nothing
        } else if fileContents[i] == '\n' {
            line_num++
        } else if (fileContents[i] == '$') {
            is_lexical_error = true
            t = token.CreateToken(token.UnexpectedCharacterError, string(fileContents[i]), "", line_num)
            tokens = append(tokens, t)
        } else if (fileContents[i] == '#') {
            is_lexical_error = true
            t = token.CreateToken(token.UnexpectedCharacterError, string(fileContents[i]), "", line_num)
            tokens = append(tokens, t)
        } else {
            is_lexical_error = true
            t = token.CreateToken(token.UnexpectedCharacterError, string(fileContents[i]), "", line_num)
            tokens = append(tokens, t)
        }
    }

    // adding EOF token at end of file
    t = token.CreateToken(token.EOF, string(""), "null", line_num)
    tokens = append(tokens, t)

    return tokens, is_lexical_error
}

func Peek() {}

func Advanced() {}

func IsAlpha(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_') 
}

func IsNumeric(c byte) bool {
    return c >= '0' && c <= '9'
}
