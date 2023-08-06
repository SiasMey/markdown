package lexer

import (
	"strings"
	"testing"
)

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

func TestScanShouldReturnTextWithNumbersLit(t *testing.T) {
	input := "Alpha129Group"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
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

func TestScanShouldReturnTextWithCommaLit(t *testing.T) {
	input := "Alpha,Comma"
	expect := input

	lex := NewScanner(strings.NewReader(input))
	_, lit := lex.Scan()

	if lit != expect {
		t.Fatalf(`Scan failed "%s" expected %s got %s`, input, expect, lit)
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

// test tooling currently doesnt do well with this it seems
func TestScanShouldReturnToken(t *testing.T) {
	tests := map[string]struct {
		input string
		want  Token
	}{
		"Eof":            {"", EOF},
		"Hash":           {"#", HASH},
		"WikiOpen":       {"[[", WIKIOPEN},
		"LeftBrc":        {"[", LEFTBRC},
		"RightBrc":       {"]", RIGHTBRC},
		"WikiClose":      {"]]", WIKICLOSE},
		"Text":           {"abc", TEXT},
		"TextSlug":       {"-b-c", TEXT},
		"TextUnderscore": {"_b_c", TEXT},
		"TextNumbers":    {"123ab", TEXT},
		"WhiteSpace":     {" ", WS},
		"WhiteSpaceTab":  {string('\t'), WS},
		"NewLineLinux":   {string('\n'), NL},
		"NewLineMac":     {string('\r'), NL},
		"NewLineDos":     {string('\r') + string('\n'), NL},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lex := NewScanner(strings.NewReader(tc.input))
			got, _ := lex.Scan()
			if got != tc.want {
				t.Fatalf(`Scan failed "%s" expected %d got %d`, tc.input, tc.want, got)
			}
		})
	}
}

func TestScanShouldReturnLiteral(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"Hash":           {"#", "#"},
		"HashContiguous": {"###", "###"},
		"Text":           {"abc", "abc"},
		"TextSlug":       {"a-b-c", "a-b-c"},
		"TextUnderscore": {"a_b_c", "a_b_c"},
		"TextNumbers":    {"925abc", "925abc"},
		"WSContiguous":   {"   ", "   "},
		"WSWithTab":      {"  	", "  	"},
		"NLLinuxSingle":  {string('\n') + string('\n'), string('\n')},
		"NLMacSingle":    {string('\r') + string('\r'), string('\r')},
		"NLDosSingle":    {string('\r') + string('\n') + "more", string('\r') + string('\n')},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lex := NewScanner(strings.NewReader(tc.input))
			_, got := lex.Scan()
			if got != tc.want {
				t.Fatalf(`Scan failed "%s" expected %v got %v`, tc.input, tc.want, got)
			}
		})
	}
}
