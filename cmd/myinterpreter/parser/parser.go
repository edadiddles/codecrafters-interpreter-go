package parser

import (
    "os"
    "fmt"
    "strings"
    "errors"
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
    Operator *token.Token
    Right Expr
}

type Parser struct {
    Tokens []*token.Token
    CurrentPosition int
}

func (p *Parser) isAtEnd() bool {
    return p.peek().Type == token.EOF
}

func (p *Parser) peek() *token.Token {
    return p.Tokens[p.CurrentPosition]
}

func (p *Parser) advance() *token.Token {
    if !p.isAtEnd() {
        p.CurrentPosition += 1
    }

    return p.previous()
}

func (p *Parser) check(tokenType token.TokenType) bool {
   if p.isAtEnd() {
        return false
    } 

    return p.peek().Type == tokenType
}

func (p *Parser) previous() *token.Token {
    return p.Tokens[p.CurrentPosition-1]
}

func (p *Parser) consume(tokenType token.TokenType, errorStr string) (*token.Token, error) {
    if p.check(tokenType) {
        return p.advance(), nil
    }

    return p.peek(), errors.New(errorStr)
}

func (p *Parser) Match(tokenTypes []token.TokenType) bool {
    for _, tokenType := range tokenTypes {
        if p.check(tokenType) {
            p.advance()
            return true
        }
    }

    return false
}

func (p *Parser) ExpressionGrammer() Expr {
    return p.EqualityGrammer()
}

func (p *Parser) EqualityGrammer() Expr {
    expr := p.ComparisonGrammer()

    for p.Match([]token.TokenType{token.BangEqual, token.EqualEqual}) {
        operator := p.previous()
        right := p.ComparisonGrammer()
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr
}

func (p *Parser) ComparisonGrammer() Expr {
    expr := p.TermGrammer()

    for p.Match([]token.TokenType{token.Greater, token.GreaterEqual, token.Less, token.LessEqual}) {
        operator := p.previous()
        right := p.TermGrammer()
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr
}

func (p *Parser) TermGrammer() Expr {
    expr := p.FactorGrammer()

    for p.Match([]token.TokenType{token.Plus, token.Minus}) {
        operator := p.previous()
        right := p.FactorGrammer()
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr
}

func (p *Parser) FactorGrammer() Expr {
    expr := p.UnaryGrammer()

    for p.Match([]token.TokenType{token.Asterisk, token.Slash}) {
        operator := p.previous()
        right := p.UnaryGrammer()
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr
}

func (p *Parser) UnaryGrammer() Expr {
    if p.Match([]token.TokenType{token.Bang, token.Minus}) {
        operator := p.previous()
        right := p.UnaryGrammer()
        return Unary{ Operator: operator, Right: right }
    }

    return p.PrimaryGrammer()
}

func (p *Parser) PrimaryGrammer() Expr {
    if p.Match([]token.TokenType{token.False,token.True,token.Nil}) {
        return Literal{ Value: p.previous() }
    } else if p.Match([]token.TokenType{token.Number, token.String}) {
        return Literal{ Value: p.previous() }
    } else if p.Match([]token.TokenType{token.LeftParen}) {
        expr := p.ExpressionGrammer()
        p.consume(token.RightParen, "Expect ')' after expression.")
        return Grouping{ Expression: expr }
    }
    
    return nil
}

func formatExpression(operator string, exprs ...Expr) string {
    str := "(" + operator
    for _, expr := range exprs {
        str = str + " "
        str = str + PrintExpression(expr)
    }
    str = str + ")"

    return str
}


func PrintExpression(expr Expr) string {
    switch e := expr.(type) {
    default:    
        return fmt.Sprintf("Unknown type: %T", e)
    case Literal:
        t := e.Value.(*token.Token)
        var val interface{}

        msg_parts := []string{ "%v" }
        if t.Type == token.Number {
            num_decimals := GetDecimalPlaces(t.Lexeme)
            val = t.Literal
            msg_parts[0] = fmt.Sprintf("%s%df", "%.", num_decimals)
        } else if t.Type == token.String {
            val = t.Literal
        } else {
            val = t.Lexeme
        }
        
        return fmt.Sprintf(strings.Join(msg_parts, ""), val)
    case Unary:
        r := e.Right
        return formatExpression(e.Operator.Lexeme, r)
    case Binary:
        l := e.Left
        r := e.Right
        return formatExpression(e.Operator.Lexeme, l, r)
    case Grouping:
        t := e.Expression
        return formatExpression("group", t)
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
   
    parser := Parser{ Tokens: tokens, CurrentPosition: 0 }
    
    for ; !parser.isAtEnd(); {
        expr := parser.ExpressionGrammer()
        expressions = append(expressions, expr)
    }
    

    for _, expr := range expressions {
        str := PrintExpression(expr)
        fmt.Fprintf(os.Stdout, "%s\n", str)
    }
}

