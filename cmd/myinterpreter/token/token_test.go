package token

import (
    "testing"
    "strconv"
)

func TestPrintTokenSimpleFloat(t *testing.T) {
    f, _ := strconv.ParseFloat("42.23", 64)
    token := Token[float64]{Type: Number, Lexeme: "42.23", Literal: f}
    token.PrintToken() 
}

func TestPrintTokenSimpleInteger(t *testing.T) {
    f, _ := strconv.ParseFloat("42", 64)
    token := Token[float64]{Type: Number, Lexeme: "42", Literal: f}
    token.PrintToken() 
}

func TestPrintTokenLeadingZerosFloat(t *testing.T) {
    f, _ := strconv.ParseFloat("42.2300", 64)
    token := Token[float64]{Type: Number, Lexeme: "42.2300", Literal: f}
    token.PrintToken() 
}

func TestPrintTokenTrailingZerosFloat(t *testing.T) {
    f, _ := strconv.ParseFloat("42.0023", 64)
    token := Token[float64]{Type: Number, Lexeme: "42.0023", Literal: f}
    token.PrintToken() 
}

func TestPrintTokenString(t *testing.T) {
    token := Token[string]{Type: String, Lexeme: "this is a string", Literal: "this is a string"}
    token.PrintToken()
}

func TestPrintTokenErrorUnterminatedString(t *testing.T) {
    token := Token[string]{ Type: Error, Lexeme: string('c'), Literal: "Unterminated String", Line: 2 }
    token.PrintToken()
}

func TestPrintTokenErrorUnexpectedToken(t *testing.T) {
    token := Token[string]{ Type: Error, Lexeme: string('$'), Literal: "Unexpected character", Line: 3 }
    token.PrintToken()
}
