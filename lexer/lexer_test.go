package lexer_test

import (
	"testing"

	"github.com/0xSherlokMo/TurtleLang/lexer"
	"github.com/0xSherlokMo/TurtleLang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tt := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)
	for i, tc := range tt {
		tok := l.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tc.expectedLiteral, tok.Literal)
		}
	}
}
