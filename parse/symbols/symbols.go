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
	WIKILINK SymbolType = "WikiLink"
)

type Symbols struct {
	Title     Symbol
	WikiLinks []Symbol
}

type Parser struct {
	s *lexer.Scanner
}

func NewParser(s string) *Parser {
	return &Parser{lexer.NewScanner(strings.NewReader(s))}
}

func Parse(input string) (Symbols, error) {
	parser := NewParser(input)
	wikiLinks := []Symbol{}

	var title Symbol

	for {
		sym, err := parser.nextSymbol()
		if err != nil {
			break
		} else if sym.Type == HEADING1 {
			title = sym
		} else if sym.Type == WIKILINK {
			wikiLinks = append(wikiLinks, sym)
		}
	}

	res := Symbols{Title: title, WikiLinks: wikiLinks}
	return res, nil
}

func (p *Parser) nextSymbol() (Symbol, error) {
	tk := p.s.Scan()

	switch tk.TokenType {
	case lexer.HASH:
		return p.parseHashStart(tk)
	case lexer.LEFTBRK:
		return p.parseLink(tk)
	}

	return Symbol{}, errors.New("Nothing left to parse")
}

func (p *Parser) parseLink(start lexer.Token) (Symbol, error) {
	charStart := start.Column
	charEnd := start.Column + start.Length
	lineNr := start.LineNr
	lit := start.Lit
	val := ""
	pairs := 1

	for {
		if tk := p.s.Scan(); tk.TokenType == lexer.EOF {
			break
		} else if tk.TokenType == lexer.LEFTBRK {
			lit += tk.Lit
			charEnd += tk.Length
			pairs += 1
		} else if tk.TokenType == lexer.TEXT {
			lit += tk.Lit
			val += tk.Lit
			charEnd += tk.Length
		} else if tk.TokenType == lexer.RIGHTBRK {
			lit += tk.Lit
			pairs -= 1
			charEnd += tk.Length
			if pairs < 1 {
				break
			}
		}
	}

	return Symbol{
		Type:      WIKILINK,
		Lit:       lit,
		Value:     val,
		LineNo:    lineNr,
		CharStart: charStart,
		CharEnd:   charEnd,
	}, nil
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
