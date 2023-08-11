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

// func TestSymbolsShouldReturnTitleIgnoringFrontmatter(t *testing.T) {
// 	input := `---
// 	tags:
// 	  - tag1
// ---
// # Title`
// 	expected := "Title"
//
// 	res, err := Parse(input)
// 	if res.Title.Value != expected {
// 		failMessageString(t, input, res.Title.Value, err, expected)
// 	}
// }

// func TestSymbolsShouldReturnTitleLoneNoIgnoringFrontmatter(t *testing.T) {
// 	input := `---
// 	tags:
// 	  - tag1
// ---
// # Title`
// 	expected := 5
//
// 	res, err := Parse(input)
// 	if res.Title.LineNo != expected {
// 		failMessageInt(t, input, res.Title.LineNo, err, expected)
// 	}
//
// }

// func TestSymbolsShouldReturnWikilinks(t *testing.T) {
// 	input := "[[wikilink]]"
// 	expected := "wikilink"
//
// 	res, err := Parse(input)
// 	if len(res.Wikilinks) != 1 {
// 		failMessageInt(t, input, len(res.Wikilinks), err, 1)
// 	}
// 	link := res.Wikilinks[0]
// 	if link.Value != expected {
// 		failMessageString(t, input, link.Value, err, expected)
// 	}
// }

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

func failMessageInt(t *testing.T, input string, result int, err error, expected int) {
	t.Fatalf(`Parse("%s") = %d, %v, expected %d`, input, result, err, expected)
}
