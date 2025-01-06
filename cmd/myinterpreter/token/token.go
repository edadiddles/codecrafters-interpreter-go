package token

import (
    "fmt"
    "os"
    "strings"
)

type TokenType int

type Token[T any] struct {
    Type TokenType
    Lexeme string
    Literal T
    Line int
}

const (
    // Single Character Tokens
    LeftParen TokenType = iota
    RightParen 
    LeftBrace
    RightBrace
    Period
    Asterisk
    Comma
    Plus
    Minus
    Semicolon
    Slash

    // One or Two Character Tokens
    Bang
    BangEqual
    Equal
    EqualEqual
    Less
    LessEqual
    Greater
    GreaterEqual

    // Keywords
    And
    Class
    Else
    False
    For
    Fun
    If
    Nil
    Or
    Print
    Super
    This
    True
    Var
    While
    Return

    // Literals
    Identifier
    String
    Number

    // End of File
    EOF

    // Error
    Error
)

var tokenTypeName = map[TokenType]string {
    // Single Character Tokens
    LeftParen: "LEFT_PAREN",
    RightParen: "RIGHT_PAREN",
    LeftBrace: "LEFT_BRACE",
    RightBrace: "RIGHT_BRACE",
    Period: "DOT",
    Asterisk: "STAR",
    Comma: "COMMA",
    Plus: "PLUS",
    Minus: "MINUS",
    Semicolon: "SEMICOLON",
    Slash: "SLASH",

    // One or Two Character Tokens
    Bang: "BANG",
    BangEqual: "BANG_EQUAL",
    Equal: "EQUAL",
    EqualEqual: "EQUAL_EQUAL",
    Less: "LESS",
    LessEqual: "LESS_EQUAL",
    Greater: "GREATER",
    GreaterEqual: "GREATER_EQUAL",

    // Keywords
    And: "AND",
    Class: "CLASS",
    Else: "ELSE",
    False: "FALSE",
    For: "FOR",
    Fun: "FUN",
    If: "IF",
    Nil: "NIL",
    Or: "OR",
    Print: "PRINT",
    Super: "SUPER",
    This: "THIS",
    True: "TRUE",
    Var: "VAR",
    While: "WHILE",

    // Literals
    Identifier: "IDENTIFIER",
    String: "STRING",
    Number: "NUMBER",

    // End of File
    EOF: "EOF",

    // Error
    Error: "ERROR",
}

func GetDecimalPlaces(floatStr string) int {
    num_decimals := 0
    for i := len(floatStr)-1; i >= 0; i-- {
        if floatStr[i] == '.' {
            break
        } else if num_decimals != 0 || floatStr[i] != '0' {
            num_decimals++
        }
    }

    if num_decimals == 0 {
        num_decimals = 1
    }

    return num_decimals
}

func (token *Token[T]) PrintToken() {
    if token.Type == Error {
        fmt.Fprintf(os.Stderr, "[line %d] Error: %v %s\n", token.Line, token.Literal, token.Lexeme)
        return 
    }

    msg_parts := []string{ "%s","%s","%s","\n" }
    if token.Type == Number {
        num_decimals := GetDecimalPlaces(token.Lexeme)
        msg_parts[2] = fmt.Sprintf("%s%df", "%.", num_decimals)
    } else if token.Type == String {
        msg_parts[1] = "\"%s\""
    }    

    fmt.Printf(strings.Join(msg_parts, " "), tokenTypeName[token.Type], token.Lexeme, token.Literal)
}
