package symbols

import (
	"reflect"
	"testing"
)

func TestParseShouldReturnTitle(t *testing.T) {
	input := "# Title"
	expected := "Title"

	res, err := Parse(input)
	if res.Title.Value != expected {
		failMessageString(t, input, res.Title.Value, err, expected)
	}
}

func TestParseShouldReturnTitleType(t *testing.T) {
	input := "# Title"
	expected := HEADING1

	res, err := Parse(input)
	if res.Title.Type != expected {
		failMessageType(t, input, res.Title.Type, err, expected)
	}
}

func TestParseShouldReturnTitleCharStart(t *testing.T) {
	input := "# Another Title"
	expected := 1

	res, err := Parse(input)
	if res.Title.CharStart != expected {
		failMessageInt(t, input, res.Title.CharStart, err, expected)
	}
}

func TestParseShouldReturnTitleCharEnd(t *testing.T) {
	input := "# New Title"
	expected := len(input) + 1

	res, err := Parse(input)
	if res.Title.CharEnd != expected {
		failMessageInt(t, input, res.Title.CharEnd, err, expected)
	}
}

func TestParseShouldReturnTitleLineNo(t *testing.T) {
	input := "# Title"
	expected := 0

	res, err := Parse(input)
	if res.Title.LineNo != expected {
		failMessageInt(t, input, res.Title.LineNo, err, expected)
	}
}

func TestParseShouldReturnWikilinkLit(t *testing.T) {
	input := "[[test]]"
	expected := "[[test]]"

	res, err := Parse(input)
	if res.WikiLinks[0].Lit != expected {
		failMessageString(t, input, res.WikiLinks[0].Lit, err, expected)
	}
}

func TestParseShouldReturnWikilinkValue(t *testing.T) {
	input := "[[test]]"
	expected := "test"

	res, err := Parse(input)
	if res.WikiLinks[0].Value != expected {
		failMessageString(t, input, res.WikiLinks[0].Value, err, expected)
	}
}

func TestParseShouldReturnWikilinkType(t *testing.T) {
	input := "[[test]]"
	expected := WIKILINK

	res, err := Parse(input)
	if res.WikiLinks[0].Type != expected {
		failMessageType(t, input, res.WikiLinks[0].Type, err, expected)
	}
}

func TestParseShouldReturnWikilinkCharStart(t *testing.T) {
	input := "[[test]]"
	expected := 1

	res, err := Parse(input)
	if res.WikiLinks[0].CharStart != expected {
		failMessageInt(t, input, res.WikiLinks[0].CharStart, err, expected)
	}
}

func TestParseShouldReturnWikilinkCharEnd(t *testing.T) {
	input := "[[test]]"
	expected := len(input) + 1

	res, err := Parse(input)
	if res.WikiLinks[0].CharEnd != expected {
		failMessageInt(t, input, res.WikiLinks[0].CharEnd, err, expected)
	}
}

func TestParseShouldReturnWikilinkLineNo(t *testing.T) {
	input := "[[test]]"
	expected := 0

	res, err := Parse(input)
	if res.WikiLinks[0].LineNo != expected {
		failMessageInt(t, input, res.WikiLinks[0].LineNo, err, expected)
	}
}

func TestParseShouldReturnTwoWikilinks(t *testing.T) {
	input := "[[test]][[test2]]"
	expected := 2

	res, err := Parse(input)
	if len(res.WikiLinks) != expected {
		failMessageInt(t, input, len(res.WikiLinks), err, expected)
	}
}

func TestParseShouldReturnLink(t *testing.T) {
	input := "[test link](http://test.com)"
	expected := LINK

	res, err := Parse(input)
	check := res.Links[0]

	if check.Type != expected {
		failMessageType(t, input, check.Type, err, expected)
	}
}

func TestParseShouldReturnLinkValue(t *testing.T) {
	input := "[test link](http://test.com)"
	expected := "http://test.com"

	res, err := Parse(input)
	check := res.Links[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestParseShouldResturnLinkStartChar(t *testing.T) {
	input := "[test link](http://trash.com)"
	expected := 1

	res, err := Parse(input)
	check := res.Links[0]

	if check.CharStart != expected {
		failMessageInt(t, input, check.CharStart, err, expected)
	}
}

func TestParseShouldResturnLinkEndChar(t *testing.T) {
	input := "[test link](http://trash.com)"
	expected := len(input) + 1

	res, err := Parse(input)
	check := res.Links[0]

	if check.CharEnd != expected {
		failMessageInt(t, input, check.CharEnd, err, expected)
	}
}

func TestParseShouldResturnLinkLineNr(t *testing.T) {
	input := "[test link](http://trash.com)"
	expected := 0

	res, err := Parse(input)
	check := res.Links[0]

	if check.LineNo != expected {
		failMessageInt(t, input, check.LineNo, err, expected)
	}
}

func TestParseShouldReturnTagLit(t *testing.T) {
	input := "#[[test]]"
	expected := "#[[test]]"

	res, err := Parse(input)
	check := res.Tags[0]

	if check.Lit != expected {
		failMessageString(t, input, check.Lit, err, expected)
	}
}

func TestParseShouldReturnTagValue(t *testing.T) {
	input := "#[[test]]"
	expected := "test"

	res, err := Parse(input)
	check := res.Tags[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestParseShouldReturnTagStartChar(t *testing.T) {
	input := "#[[test]]"
	expected := 1

	res, err := Parse(input)
	check := res.Tags[0]

	if check.CharStart != expected {
		failMessageInt(t, input, check.CharStart, err, expected)
	}
}

func TestParseShouldReturnTagEndChar(t *testing.T) {
	input := "#[[test]]"
	expected := len(input) + 1

	res, err := Parse(input)
	check := res.Tags[0]

	if check.CharEnd != expected {
		failMessageInt(t, input, check.CharEnd, err, expected)
	}
}

func TestParseShouldReturnTagLineNr(t *testing.T) {
	input := "#[[test]]"
	expected := 0

	res, err := Parse(input)
	check := res.Tags[0]

	if check.LineNo != expected {
		failMessageInt(t, input, check.LineNo, err, expected)
	}
}

func TestParseShouldReturnHeadersValue(t *testing.T) {
	input := "## test heading 2"
	expected := "test heading 2"

	res, err := Parse(input)
	check := res.Headers[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestParseShouldReturnCombinationsTitle(t *testing.T) {
	input := `# test header
[Http link](http://test.com) [[arst1234]]
	#[[test-tag]]
`

	res, err := Parse(input)

	if res.Title.Value != "test header" {
		failMessageString(t, input, res.Title.Value, err, "test header")
	}
}

func TestParseShouldReturnCombinationsLink(t *testing.T) {
	input := `# test header
[Http link](http://test.com) [[arst1234]]
	#[[test-tag]]
`
	expected := "http://test.com"

	res, err := Parse(input)
	check := res.Links[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestParseShouldReturnCombinationsWikilink(t *testing.T) {
	input := `# test header
[Http link](http://test.com) [[arst1234]]
	#[[test-tag]]
`
	expected := "arst1234"

	res, err := Parse(input)
	check := res.WikiLinks[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
	}
}

func TestParseShouldReturnCombinationsTag(t *testing.T) {
	input := `# test header
[Http link](http://test.com) [[arst1234]]
	#[[test-tag]]
`
	expected := "test-tag"

	res, err := Parse(input)
	check := res.Tags[0]

	if check.Value != expected {
		failMessageString(t, input, check.Value, err, expected)
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
