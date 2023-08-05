package symbols

import (
	"reflect"
	"testing"
)

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

func TestSymbolsShouldReturnTitle(t *testing.T) {
	input := "# Title"
	expected := "Title"

	res, err := Extract(input)
	if res.Title.Value != expected {
		t.Fatalf(`Extract(%s) = %q, %v, expected %s`, input, res.Title.Value, err, expected)
	}
}

func TestSymbolsShouldReturnWikilinks(t *testing.T) {
	input := "[[wikilink]]"
	expected := "wikilink"

	res, err := Extract(input)
	if len(res.Wikilinks) < 1 {
		t.Fatalf(`Extract(%s) Expected at least one wikilink: %q, %v`, input, res.Wikilinks, err)
	}

	link := res.Wikilinks[0]
	if link.Value != expected {
		t.Fatalf(`Extract(%s) = %q, %v, expected %s`, input, res.Wikilinks, err, expected)
	}
}
