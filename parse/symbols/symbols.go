package symbols

type Symbol struct {
	Value string
}

type Symbols struct {
	Title     Symbol
	Wikilinks []Symbol
}

func Extract(input string) (Symbols, error) {
	return Symbols{Title: Symbol{Value: "Title"}, Wikilinks: []Symbol{{"wikilink"}}}, nil
}
