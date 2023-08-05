package symbols

type Symbols struct {
  Title string
  Wikilinks []string
}

func Extract(input string) (Symbols, error) {
	return Symbols{Title: "Title", Wikilinks: []string{"wikilink"}}, nil
}
