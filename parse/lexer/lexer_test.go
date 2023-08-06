package lexer

import (
	"strings"
	"testing"
)

func TestScanShouldReturnHash(t *testing.T) {
	input := "#"
	expect := HASH

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d` , input, expect, tok)
	}
}

func TestScanShouldReturnLitHash(t *testing.T) {
	input := "#"
	expect := "#"

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s` , input, expect, lit)
	}
}

func TestScanShouldReturnWikiOpen(t *testing.T) {
	input := "[["
	expect := WIKIOPEN

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d` , input, expect, tok)
	}
}

func TestScanShouldReturnWikiClose(t *testing.T) {
	input := "]]"
	expect := WIKICLOSE

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d` , input, expect, tok)
	}
}

func TestScanShouldReturnLeftBrc(t *testing.T) {
	input := "["
	expect := LEFTBRC

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d` , input, expect, tok)
	}
}

func TestScanShouldReturnRightBrc(t *testing.T) {
	input := "]"
	expect := RIGHTBRC

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d` , input, expect, tok)
	}
}
