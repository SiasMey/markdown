package symbols

import (
	"reflect"
	"testing"
)

func TestSymbolsShouldReturnTitle(t *testing.T) {
	input := "# Title"
	expected := "Title"

	res, err := Parse(input)
	if res.Title.Value != expected {
		failMessageString(t, input, res.Title.Value, err, expected)
	}
}

func TestSymbolsShouldReturnTitleType(t *testing.T) {
	input := "# Title"
	expected := HEADING1

	res, err := Parse(input)
	if res.Title.Type != expected {
		failMessageType(t, input, res.Title.Type, err, expected)
	}
}

func TestSymbolsShouldReturnTitleCharStart(t *testing.T) {
	input := "# Another Title"
	expected := 1

	res, err := Parse(input)
	if res.Title.CharStart != expected {
		failMessageInt(t, input, res.Title.CharStart, err, expected)
	}
}

func TestSymbolsShouldReturnTitleCharEnd(t *testing.T) {
	input := "# New Title"
	expected := len(input) + 1

	res, err := Parse(input)
	if res.Title.CharEnd != expected {
		failMessageInt(t, input, res.Title.CharEnd, err, expected)
	}
}

func TestSymbolsShouldReturnTitleLineNo(t *testing.T) {
	input := "# Title"
	expected := 0

	res, err := Parse(input)
	if res.Title.LineNo != expected {
		failMessageInt(t, input, res.Title.LineNo, err, expected)
	}
}

func TestSymbolShouldReturnWikilinkLit(t *testing.T) {
	input := "[[test]]"
	expected := "[[test]]"

	res, err := Parse(input)
	if res.WikiLinks[0].Lit != expected {
		failMessageString(t, input, res.WikiLinks[0].Lit, err, expected)
	}
}

func TestSymbolShouldReturnWikilinkValue(t *testing.T) {
	input := "[[test]]"
	expected := "test"

	res, err := Parse(input)
	if res.WikiLinks[0].Value != expected {
		failMessageString(t, input, res.WikiLinks[0].Value, err, expected)
	}
}

func TestSymbolsShouldReturnWikilinkType(t *testing.T) {
	input := "[[test]]"
	expected := WIKILINK

	res, err := Parse(input)
	if res.WikiLinks[0].Type != expected {
		failMessageType(t, input, res.WikiLinks[0].Type, err, expected)
	}
}

func TestSymbolsShouldReturnWikilinkCharStart(t *testing.T) {
	input := "[[test]]"
	expected := 1

	res, err := Parse(input)
	if res.WikiLinks[0].CharStart != expected {
		failMessageInt(t, input, res.WikiLinks[0].CharStart, err, expected)
	}
}

func TestSymbolsShouldReturnWikilinkCharEnd(t *testing.T) {
	input := "[[test]]"
	expected := len(input) + 1

	res, err := Parse(input)
	if res.WikiLinks[0].CharEnd != expected {
		failMessageInt(t, input, res.WikiLinks[0].CharEnd, err, expected)
	}
}

func TestSymbolsShouldReturnWikilinkLineNo(t *testing.T) {
	input := "[[test]]"
	expected := 0

	res, err := Parse(input)
	if res.WikiLinks[0].LineNo != expected {
		failMessageInt(t, input, res.WikiLinks[0].LineNo, err, expected)
	}
}

func TestSymbolsShouldReturnTwoWikilinks(t *testing.T) {
	input := "[[test]][[test2]]"
	expected := 2

	res, err := Parse(input)
	if len(res.WikiLinks) != expected {
		failMessageInt(t, input, len(res.WikiLinks), err, expected)
	}
}

func TestSymbolsShouldReturnLink(t *testing.T) {
	input := "[test link](http://test.com)"
	expected := LINK

	res, err := Parse(input)
	check := res.Links[0]

	if check.Type != expected {
		failMessageType(t, input, check.Type, err, expected)
	}
}

func TestSymbolsShouldReturnLinkValue(t *testing.T) {
	input := "[test link](http://test.com)"
	expected := "http://test.com"

	res, err := Parse(input)
	check := res.Links[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestSymbolsShouldResturnLinkStartChar(t *testing.T) {
	input := "[test link](http://trash.com)"
	expected := 1

	res, err := Parse(input)
	check := res.Links[0]

	if check.CharStart != expected {
		failMessageInt(t, input, check.CharStart, err, expected)
	}
}

func TestSymbolsShouldResturnLinkEndChar(t *testing.T) {
	input := "[test link](http://trash.com)"
	expected := len(input) + 1

	res, err := Parse(input)
	check := res.Links[0]

	if check.CharEnd != expected {
		failMessageInt(t, input, check.CharEnd, err, expected)
	}
}

func TestSymbolsShouldResturnLinkLineNr(t *testing.T) {
	input := "[test link](http://trash.com)"
	expected := 0

	res, err := Parse(input)
	check := res.Links[0]

	if check.LineNo != expected {
		failMessageInt(t, input, check.LineNo, err, expected)
	}
}

func TestSymbolsShouldReturnTagLit(t *testing.T) {
	input := "#[[test]]"
	expected := "#[[test]]"

	res, err := Parse(input)
	check := res.Tags[0]

	if check.Lit != expected {
		failMessageString(t, input, check.Lit, err, expected)
	}
}

func TestSymbolsShouldReturnTagValue(t *testing.T) {
	input := "#[[test]]"
	expected := "test"

	res, err := Parse(input)
	check := res.Tags[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestSymbolsShouldResturnTagStartChar(t *testing.T) {
	input := "#[[test]]"
	expected := 1

	res, err := Parse(input)
	check := res.Tags[0]

	if check.CharStart != expected {
		failMessageInt(t, input, check.CharStart, err, expected)
	}
}

func TestSymbolsShouldResturnTagEndChar(t *testing.T) {
	input := "#[[test]]"
	expected := len(input) + 1

	res, err := Parse(input)
	check := res.Tags[0]

	if check.CharEnd != expected {
		failMessageInt(t, input, check.CharEnd, err, expected)
	}
}

func TestSymbolsShouldResturnTagLineNr(t *testing.T) {
	input := "#[[test]]"
	expected := 0

	res, err := Parse(input)
	check := res.Tags[0]

	if check.LineNo != expected {
		failMessageInt(t, input, check.LineNo, err, expected)
	}
}


func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type to itemExists")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func failMessageString(t *testing.T, input string, result string, err error, expected string) {
	t.Fatalf(`Parse("%s") = %q, %v, expected %s`, input, result, err, expected)
}

func failMessageType(t *testing.T, input string, result SymbolType, err error, expected SymbolType) {
	t.Fatalf(`Parse("%s") = %q, %v, expected %s`, input, result, err, expected)
}

func failMessageInt(t *testing.T, input string, result int, err error, expected int) {
	t.Fatalf(`Parse("%s") = %d, %v, expected %d`, input, result, err, expected)
}
