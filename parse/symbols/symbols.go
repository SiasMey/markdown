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
	HEADING2 SymbolType = "Heading2"
	WIKILINK SymbolType = "WikiLink"
	LINK     SymbolType = "Link"
	TAG      SymbolType = "Tag"
	OTHER    SymbolType = "Other"
)

type Symbols struct {
	Title     Symbol
	WikiLinks []Symbol
	Links     []Symbol
	Tags      []Symbol
	Headers   []Symbol
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
	links := []Symbol{}
	tags := []Symbol{}
	headers := []Symbol{}

	var title Symbol

	for {
		sym, err := parser.nextSymbol()
		if err != nil {
			break
		} else if sym.Type == HEADING1 {
			title = sym
		} else if sym.Type == HEADING2 {
			headers = append(headers, sym)
		} else if sym.Type == WIKILINK {
			wikiLinks = append(wikiLinks, sym)
		} else if sym.Type == LINK {
			links = append(links, sym)
		} else if sym.Type == TAG {
			tags = append(tags, sym)
		}
	}

	res := Symbols{
		Title:     title,
		WikiLinks: wikiLinks,
		Links:     links,
		Tags:      tags,
		Headers:   headers,
	}
	return res, nil
}

func (p *Parser) nextSymbol() (Symbol, error) {
	tk := p.s.Scan()

	switch tk.TokenType {
	case lexer.HASH:
		return p.parseHashStart(tk)
	case lexer.LEFTBRK:
		return p.parseLink(tk)
	case lexer.EOF:
		return Symbol{}, errors.New("Nothing left to parse")
	default:
		return Symbol{
			Type:      OTHER,
			Lit:       tk.Lit,
			Value:     tk.Lit,
			LineNo:    tk.LineNr,
			CharStart: tk.Column,
			CharEnd:   tk.Column + tk.Length,
		}, nil
	}
}

func (p *Parser) parseLink(start lexer.Token) (Symbol, error) {
	charStart := start.Column
	charEnd := start.Column + start.Length
	lineNr := start.LineNr
	lit := start.Lit
	linkType := LINK
	val := ""
	pairs := 1

	for {
		if tk := p.s.Scan(); tk.TokenType == lexer.EOF {
			break
		} else if tk.TokenType == lexer.LEFTBRK {
			lit += tk.Lit
			charEnd += tk.Length
			pairs += 1
			linkType = WIKILINK
		} else if tk.TokenType == lexer.TEXT {
			lit += tk.Lit
			charEnd += tk.Length
			val += tk.Lit
		} else if tk.TokenType == lexer.WS {
			lit += tk.Lit
			charEnd += tk.Length
			val += tk.Lit
		} else if tk.TokenType == lexer.ILLEGAL {
			lit += tk.Lit
			charEnd += tk.Length
			val += tk.Lit
		} else if tk.TokenType == lexer.RIGHTBRK {
			lit += tk.Lit
			pairs -= 1
			charEnd += tk.Length

			if pairs < 1 {
				if linkType == WIKILINK {
					break
				}
			}
		} else if tk.TokenType == lexer.LEFTPRN {
			lit += tk.Lit
			val = ""
			charEnd += tk.Length
		} else if tk.TokenType == lexer.RIGHTPRN {
			lit += tk.Lit
			charEnd += tk.Length
			break
		}
	}

	return Symbol{
		Type:      linkType,
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
	hashType := HEADING1
	gotTrailingWs := false
	scopes := 0

	if start.Length == 2 {
		hashType = HEADING2
	}

	for {
		if tk := p.s.Scan(); tk.TokenType == lexer.EOF {
			break
		} else if tk.TokenType == lexer.NL {
			break
		} else if start.Lit == "#" && tk.TokenType == lexer.LEFTBRK {
			hashType = TAG
			lit += tk.Lit
			charEnd += tk.Length

			scopes += 1
		} else if tk.TokenType == lexer.RIGHTBRK {
			scopes -= 1
			lit += tk.Lit
			charEnd += tk.Length

			if hashType == TAG && scopes < 1 {
				break
			}
		} else if tk.TokenType == lexer.WS {
			if gotTrailingWs {
				lit += tk.Lit
				charEnd += tk.Length
				val += tk.Lit
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
		Type:      hashType,
		Lit:       lit,
		Value:     val,
		LineNo:    lineNr,
		CharStart: charStart,
		CharEnd:   charEnd,
	}, nil
}
