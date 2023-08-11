package symbols

import (
	"errors"
	"strings"

	"github.com/siasmey/markdown/parse/lexer"
)

type Symbol struct {
	Lit       string
	Value     string
	CharStart int
	CharEnd   int
	LineNo    int
	Type      SymbolType
}

type SymbolType string

const (
	HEADING1 SymbolType = "Heading1"
)

type Symbols struct {
	Title Symbol
}

type Parser struct {
	s *lexer.Scanner
}

func (p *Parser) parseHashStart(start lexer.Token) (Symbol, error) {
	charStart := start.Column
	charEnd := start.Column + start.Length
	lineNr := start.LineNr
	lit := start.Lit
	val := ""
	gotTrailingWs := false

	for {
		if tk := p.s.Scan(); tk.TokenType == lexer.EOF {
			break
		} else if tk.TokenType == lexer.NL {
			break
		} else if tk.TokenType == lexer.WS {
			if gotTrailingWs {
				lit += tk.Lit
				val += tk.Lit
				charEnd += tk.Length
			} else {
				gotTrailingWs = true
				lit += tk.Lit
				charEnd += tk.Length
			}
		} else {
			lit += tk.Lit
			val += tk.Lit
			charEnd += tk.Length
		}
	}

	return Symbol{
		Type:      HEADING1,
		Lit:       lit,
		Value:     val,
		LineNo:    lineNr,
		CharStart: charStart,
		CharEnd:   charEnd,
	}, nil
}

func (p *Parser) nextSymbol() (Symbol, error) {
	tk := p.s.Scan()

	switch tk.TokenType {
	case lexer.HASH:
		return p.parseHashStart(tk)
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
