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

func TestSymbolsShouldReturnTitleCharStart(t *testing.T) {
	input := "# Title"
	expected := 0

	res, err := Extract(input)
	if res.Title.CharStart != expected {
		t.Fatalf(`Extract(%s) = %d, %v, expected %d`, input, res.Title.CharStart, err, expected)
	}
}

func failMessageString(t *testing.T, input string, result string, err error, expected string) {
	t.Fatalf(`Extract(%s) = %q, %v, expected %s`, input, result, err, expected)
}

func failMessageInt(t *testing.T, input string, result int, err error, expected int) {
	t.Fatalf(`Extract(%s) = %d, %v, expected %d`, input, result, err, expected)
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
