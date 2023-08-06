package lexer

import (
	"strings"
	"testing"
)

func TestScanShouldReturnEof(t *testing.T) {
	input := ""
	expect := EOF

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnHash(t *testing.T) {
	input := "#"
	expect := HASH

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnLitHash(t *testing.T) {
	input := "#"
	expect := "#"

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnCongtiguousHashLit(t *testing.T) {
	input := "###"
	expect := "###"

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnWikiOpen(t *testing.T) {
	input := "[["
	expect := WIKIOPEN

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnWikiClose(t *testing.T) {
	input := "]]"
	expect := WIKICLOSE

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnLeftBrc(t *testing.T) {
	input := "["
	expect := LEFTBRC

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnRightBrc(t *testing.T) {
	input := "]"
	expect := RIGHTBRC

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnText(t *testing.T) {
	input := "AlphaGroup"
	expect := TEXT

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnTextLit(t *testing.T) {
	input := "AlphaGroup"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnTextSlugLit(t *testing.T) {
	input := "text-slug-test"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnTextUnderCaseLit(t *testing.T) {
	input := "text_slug_test"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnTextWithNumbers(t *testing.T) {
	input := "Alpha129Group"
	expect := TEXT

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnTextWithNumbersLit(t *testing.T) {
	input := "Alpha129Group"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnTextWithFullstop(t *testing.T) {
	input := "Alpha.Fullstop"
	expect := TEXT

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnTextWithFullstopLit(t *testing.T) {
	input := "Alpha.Fullstop"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnTextWithComma(t *testing.T) {
	input := "Alpha,Comma"
	expect := TEXT

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnTextWithCommaLit(t *testing.T) {
	input := "Alpha,Comma"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnWS(t *testing.T) {
	input := " "
	expect := WS

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnContiguosWSLit(t *testing.T) {
	input := "      "
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
	}
}

func TestScanShouldReturnWSTab(t *testing.T) {
	input := "	"
	expect := WS

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnNLforLF(t *testing.T) {
	input := string('\n')
	expect := NL

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnNLforCR(t *testing.T) {
	input := string('\r')
	expect := NL

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnNLforCRLF(t *testing.T) {
	input := string('\r') + string('\n')
	expect := NL

	lex := NewScanner(strings.NewReader(input))
	tok, _ := lex.Scan()

	if tok != expect {
		t.Fatalf(`Scan failed "%s" expected %d got %d`, input, expect, tok)
	}
}

func TestScanShouldReturnSingleNLLitLinux(t *testing.T) {
	input := string('\n') + string('\n')
	expect := string('\n')

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected "%s" got "%s"`, input, expect, lit)
	}
}

func TestScanShouldReturnSingleNLLitMac(t *testing.T) {
	input := string('\r') + string('\r')
	expect := string('\r')

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected "%s" got "%s"`, input, expect, lit)
	}
}

func TestScanShouldReturnSingleNLLitWin(t *testing.T) {
	input := string('\r') + string('\n')
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected "%s" got "%s"`, input, expect, lit)
	}
}
