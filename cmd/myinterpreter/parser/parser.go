package parser

import (
    "os"
    "fmt"
    "strings"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)

type Expr interface {}

type Binary struct {
    Left Expr
    Right Expr
    Operator *token.Token
}

type Grouping struct {
    Expression Expr
}

type Literal struct {
    Value interface{}
}

type Unary struct {
    Operator token.Token
    right Expr
}

func peek() token.Token {
    return token.Token{}
}

func advance() token.Token {
    return token.Token{}
}

func check(tokenType token.TokenType) bool {
    if tokenType == token.EOF {
        return false
    }

    return peek().Type == tokenType
}

func previous() token.Token {
    return token.Token{}
}

func Match(tokenTypes []token.TokenType) bool {
    for _, tokenType := range tokenTypes {
        if check(tokenType) {
            advance()
            return true
        }
    }

    return false
}

//func Expression() Expr {
//    return Equality()
//}

//func Equality() Expr {
//    expr := Comparison()

//    for Match([]token.TokenType{token.BangEqual, token.EqualEqual}) {
//        operator := previous()
//        right := Comparison()
//        expr = Binary{ Left: expr, Right: right, Operator: *operator }
//    }

//    return expr
//}

//func Comparison() Expr {
//    return Binary{}
//}

func Addition() {
}

func Multiplication() {
}

func PrintExpression(expr Expr) {
    switch e := expr.(type) {
    default:    
        fmt.Println("default case")
    case Literal:
        t := e.Value.(*token.Token)
        fmt.Fprintf(os.Stdout, "%v\n", t.Lexeme)
    case Binary:
        l := e.Left.(*token.Token)
        r := e.Right.(*token.Token)

        msg_parts := []string{ "(%s","%s","%s)\n" }
        if l.Type == token.Number {
            num_decimals := GetDecimalPlaces(l.Lexeme)
            msg_parts[1] = fmt.Sprintf("%s%df", "%.", num_decimals)
        } else if l.Type == token.String {
            msg_parts[1] = "\"%s\""
        }    
        if r.Type == token.Number {
            num_decimals := GetDecimalPlaces(r.Lexeme)
            msg_parts[2] = fmt.Sprintf("%s%df)\n", "%.", num_decimals)
        } else if l.Type == token.String {
            msg_parts[2] = "\"%s\")\n"
        }
        fmt.Fprintf(os.Stdout, strings.Join(msg_parts, ", "), e.Operator.Lexeme, l.Literal, r.Literal)
    }
}

func GetDecimalPlaces(floatStr string) int {
    num_decimals := 0
    has_decimal := false
    for i:=len(floatStr)-1; i >= 0; i-- {
        if floatStr[i] == '.' {
            has_decimal = true
            break
        } else if num_decimals != 0 || floatStr[i] != '0' {
            num_decimals++
        }
    }

    if num_decimals == 0 || !has_decimal {
        num_decimals = 1
    }

    return num_decimals
}

func Parse(tokens []*token.Token) {
    expressions := make([]Expr, 0) 
    for i:=0; i < len(tokens); i++ {
        currType := tokens[i].Type
        if currType == token.True || currType == token.False || currType == token.Nil {
            expressions = append(expressions, Literal{ Value: tokens[i] })
        } else if currType == token.Plus || currType == token.Minus {
            expressions = append(expressions, Binary{Left: tokens[i-1], Right: tokens[i+1], Operator: tokens[i]})
        } else if currType == token.Asterisk || currType == token.Slash {
            expressions = append(expressions, Binary{Left: tokens[i-1], Right: tokens[i+1], Operator: tokens[i]})
        }
    }

    for _, expr := range expressions {
        PrintExpression(expr)
    }
}

