package evaluator

import (
    "fmt"
    "errors"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
    "github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
)

type Eval interface{}

func isTruthy(eval Eval) bool { 
    if eval == "nil" {
        return false
    }

    v, ok := eval.(bool)
    if ok {
        return v
    }

    return true
}

func isEqual(e1 Eval, e2 Eval) bool {
    return e1 == e2
}


func EvaluateExpression(expr parser.Expr) (Eval, error) {
    switch e := expr.(type) {
    default:
        return nil, nil
    case parser.Literal:
        return EvaluateLiteralExpression(e)
    case parser.Unary:
        return EvaluateUnaryExpression(e)
    case parser.Grouping:
        return EvaluateGroupingExpression(e)
    case parser.Binary:
        return EvaluateBinaryExpression(e)
    }

    // Unreachable
    return nil, errors.New("unreachable")
}

func EvaluateBinaryExpression(expr parser.Binary) (Eval, error) {
    left, err := EvaluateExpression(expr.Left)
    if err != nil {
        return nil, err
    }
    right, err := EvaluateExpression(expr.Right)
    if err != nil {
        return nil, err
    }

    if expr.Operator.Type == token.Minus {
        l := left.(float64)
        r := right.(float64)
        return l-r, nil
    } else if expr.Operator.Type == token.Slash {
        l := left.(float64)
        r := right.(float64)
        return l/r, nil
    } else if expr.Operator.Type == token.Asterisk {
        l := left.(float64)
        r := right.(float64)
        return l*r, nil
    } else if expr.Operator.Type == token.Plus {
        lf_val, lf_ok := left.(float64)
        rf_val, rf_ok := right.(float64)
        if lf_ok && rf_ok {
            return lf_val+rf_val, nil
        }

        lstr_val, lstr_ok := left.(string)
        rstr_val, rstr_ok := right.(string)
        if lstr_ok && rstr_ok {
            return lstr_val+rstr_val, nil
        }

        return nil, errors.New("Invalid type, expected float64 or string")
    } else if expr.Operator.Type == token.Greater {
        l, ok := left.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        r, ok := right.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        return l>r, nil
    } else if expr.Operator.Type == token.GreaterEqual {
        l, ok := left.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        r, ok := right.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        return l>=r, nil
    } else if expr.Operator.Type == token.Less {
        l, ok := left.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        r, ok := right.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        return l<r, nil
    } else if expr.Operator.Type == token.LessEqual {
        l, ok := left.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        r, ok := right.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        return l<=r, nil
    } else if expr.Operator.Type == token.BangEqual {
        return !isEqual(left,right), nil
    } else if expr.Operator.Type == token.EqualEqual {
        return isEqual(left,right), nil
    }

    // Unreachable
    return nil, errors.New("unreachable")
}

func EvaluateGroupingExpression(expr parser.Grouping) (Eval, error) {
    return EvaluateExpression(expr.Expression)
}

func EvaluateLiteralExpression(expr parser.Literal) (Eval, error) {
    t := expr.Value.(*token.Token)
    if t.Type == token.Number || t.Type == token.String { 
        return t.Literal, nil
    } else if t.Type == token.True {
        return true, nil
    } else if t.Type == token.False {
        return false, nil
    }

    return t.Lexeme, nil
}

func EvaluateUnaryExpression(expr parser.Unary) (Eval, error) {
    right, err := EvaluateExpression(expr.Right)
    if err != nil {
        return nil, err
    }
    if expr.Operator.Type == token.Minus {
        r, ok := right.(float64)
        if !ok {
            return nil, errors.New("Unexpected type, expected float64")
        }
        return -r, nil
    } else if expr.Operator.Type == token.Bang {
        return !isTruthy(right), nil
    }

    // Unreachable
    return nil, errors.New("unreachable")
}

func Evaluate(exprs []parser.Expr) ([]Eval, bool) {
    evals := make([]Eval, 0)
    
    hasError := false
    for i:=0; i < len(exprs); i++ {
        e, err := EvaluateExpression(exprs[i])
        if err != nil {
            hasError = true
            continue
        }
        evals = append(evals, e)
    }

    return evals, hasError
}

func PrintEval(eval Eval) string {
    return fmt.Sprintf("%v", eval)
}
