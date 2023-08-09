package lexer

import (
	"strings"
	"testing"
)

func TestScanShouldReturnToken(t *testing.T) {
	tests := map[string]struct {
		input string
		want  TokenType
	}{
		"Eof":            {"", EOF},
		"Hash":           {"#", HASH},
		"LeftBrc":        {"[", LEFTBRK},
		"RightBrc":       {"]", RIGHTBRK},
		"LeftParen":      {"(", LEFTPRN},
		"RightParen":     {")", RIGHTPRN},
		"Tick":           {"`", TICK},
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
			got := lex.Scan()
			if got.TypeType != tc.want {
				t.Fatalf(`Scan failed "%s" expected %d got %d`, tc.input, tc.want, got.TypeType)
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
			got := lex.Scan()
			if got.Lit != tc.want {
				t.Fatalf(`Scan failed "%s" expected %v got %v`, tc.input, tc.want, got.Lit)
			}
		})
	}
}

func TestScanTextLiteralShouldNotOverrun(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"LeftBrc":       {"ast[", "ast"},
		"RightBrc":      {"ast]", "ast"},
		"LeftPrn":       {"ast(", "ast"},
		"RightPrn":      {"ast)", "ast"},
		"Hash":          {"ast#", "ast"},
		"TICK":          {"ast`", "ast"},
		"WhiteSpace":    {"ast ", "ast"},
		"WhiteSpaceTab": {"ast	", "ast"},
		"NewlineNix":    {"ast" + string('\n'), "ast"},
		"NewlineMac":    {"ast" + string('\r'), "ast"},
		"NewlineDos":    {"ast" + string('\r') + string('\n'), "ast"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lex := NewScanner(strings.NewReader(tc.input))
			got := lex.Scan()
			if got.Lit != tc.want {
				t.Fatalf(`Scan failed "%s" expected %v got %v`, tc.input, tc.want, got.Lit)
			}
		})
	}
}

func TestScanWhiteSpaceShouldNotOverrun(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"SpaceLeftBrc":  {" [", " "},
		"SpaceRightBrc": {" ]", " "},
		"SpaceLeftPrn":  {" (", " "},
		"SpaceRightPrn": {" )", " "},
		"SpaceHash":     {" #", " "},
		"SpaceTICK":     {" `", " "},
		"TabLeftBrc":    {"	[", "	"},
		"TabRightBrc":   {"	]", "	"},
		"TabLeftPrn":    {"	(", "	"},
		"TabRightPrn":   {"	)", "	"},
		"TabHash":       {"	#", "	"},
		"TabTICK":       {"	`", "	"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lex := NewScanner(strings.NewReader(tc.input))
			got := lex.Scan()
			if got.Lit != tc.want {
				t.Fatalf(`Scan failed "%s" expected %v got %v`, tc.input, tc.want, got.Lit)
			}
		})
	}
}
