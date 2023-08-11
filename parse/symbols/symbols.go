package symbols

import (
	"errors"
	"strings"

	"github.com/siasmey/markdown/parse/lexer"
)

type Symbol struct {
	Value     string
	CharStart int
	CharEnd   int
	LineNo    int
}

type Symbols struct {
	Title Symbol
}

type Parser struct {
	s *lexer.Scanner
}

func (p *Parser) parseHeader(start lexer.Token) (Symbol, error) {
	charStart := start.Column
	charEnd := start.Column + start.Length
	lineNr := start.LineNr
	lit := start.Lit

	for {
		if tk := p.s.Scan(); tk.TokenType == lexer.EOF {
			break
		} else if tk.TokenType == lexer.NL {
			break
		} else {
			lit += tk.Lit
			charEnd += tk.Length
		}
	}

	return Symbol{Value: lit, LineNo: lineNr, CharStart: charStart, CharEnd: charEnd}, nil
}

func (p *Parser) nextSymbol() (Symbol, error) {
	tk := p.s.Scan()

	switch tk.TokenType {
	case lexer.HASH:
		return p.parseHeader(tk)
	}

	return Symbol{}, errors.New("Nothing left to parse")
}

func NewParser(s string) *Parser {
	return &Parser{lexer.NewScanner(strings.NewReader(s))}
}

func Parse(input string) (Symbols, error) {
	parser := NewParser(input)
	title, _ := parser.nextSymbol()

	res := Symbols{Title: title}
	return res, nil
}
