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

func (p *Parser) synchronize() {
    p.advance()

    for ; !p.isAtEnd(); {
        if p.previous().Type == token.Semicolon {
            return
        }

        t := p.peek().Type
        if t == token.Class {
            return
        } else if t == token.Fun {
            return
        } else if t == token.Var {
            return
        } else if t == token.For {
            return
        } else if t == token.If {
            return
        } else if t == token.While {
            return
        } else if t == token.Print {
            return
        } else if t == token.Return {
            return
        }

        p.advance()
    }
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

func (p *Parser) ExpressionGrammer() (Expr, error) {
    return p.EqualityGrammer()
}

func (p *Parser) EqualityGrammer() (Expr, error) {
    expr, err := p.ComparisonGrammer()
    if err != nil {
        return nil, err
    }

    for p.Match([]token.TokenType{token.BangEqual, token.EqualEqual}) {
        operator := p.previous()
        right, err := p.ComparisonGrammer()
        if err != nil {
            return nil, err
        }

        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr, nil
}

func (p *Parser) ComparisonGrammer() (Expr, error) {
    expr, err := p.TermGrammer()
    if err != nil {
        return nil, err
    }

    for p.Match([]token.TokenType{token.Greater, token.GreaterEqual, token.Less, token.LessEqual}) {
        operator := p.previous()
        right, err := p.TermGrammer()
        if err != nil {
            return nil, err
        }
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr, nil
}

func (p *Parser) TermGrammer() (Expr, error) {
    expr, err := p.FactorGrammer()
    if err != nil {
        return nil, err
    }

    for p.Match([]token.TokenType{token.Plus, token.Minus}) {
        operator := p.previous()
        right, err := p.FactorGrammer()
        if err != nil {
            return nil, err
        }
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr, nil
}

func (p *Parser) FactorGrammer() (Expr, error) {
    expr, err := p.UnaryGrammer()
    if err != nil {
        return nil, err
    }

    for p.Match([]token.TokenType{token.Asterisk, token.Slash}) {
        operator := p.previous()
        right, err := p.UnaryGrammer()
        if err != nil {
            return nil, err
        }
        expr = Binary{ Left: expr, Right: right, Operator: operator }
    }

    return expr, nil
}

func (p *Parser) UnaryGrammer() (Expr, error) {
    if p.Match([]token.TokenType{token.Bang, token.Minus}) {
        operator := p.previous()
        right, err := p.UnaryGrammer()
        if err != nil {
            return nil, err
        }
        return Unary{ Operator: operator, Right: right }, nil
    }

    return p.PrimaryGrammer()
}

func (p *Parser) PrimaryGrammer() (Expr, error) {
    if p.Match([]token.TokenType{token.False,token.True,token.Nil}) {
        return Literal{ Value: p.previous() }, nil
    } else if p.Match([]token.TokenType{token.Number, token.String}) {
        return Literal{ Value: p.previous() }, nil
    } else if p.Match([]token.TokenType{token.LeftParen}) {
        expr, err := p.ExpressionGrammer()
        if err != nil {
            return nil, err
        }
        t, err := p.consume(token.RightParen, "Expect ')' after expression.")
        if err != nil {
            if t.Type == token.EOF {
                fmt.Fprintf(os.Stderr, "[line %d] Error: At EOF\n", t.Line)
            } else {
                fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", t.Line, err)
            }
            return nil, err
        }
        return Grouping{ Expression: expr }, nil
    } else if p.Match([]token.TokenType{token.UnterminatedStringError, token.NumericError, token.UnexpectedCharacterError}) {
        fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", p.previous().Line, "Lexical Error")
        return nil, errors.New("Lexical Error")
    }

    p.advance()
    return nil, errors.New("Unknown Grammer")
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

func Parse(tokens []*token.Token) ([]Expr, bool) {
    expressions := make([]Expr, 0) 
  
    hasError := false
    parser := Parser{ Tokens: tokens, CurrentPosition: 0 } 
    for ; !parser.isAtEnd(); {
        expr, err := parser.ExpressionGrammer()
        if err != nil {
            hasError = true
            continue
        }
        expressions = append(expressions, expr)
    }

    return expressions, hasError
}

