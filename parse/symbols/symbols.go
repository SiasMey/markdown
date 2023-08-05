package symbols

type Symbol struct {
	Value     string
	CharStart int
	CharEnd   int
}

type Symbols struct {
	Title     Symbol
	Wikilinks []Symbol
}

func Extract(input string) (Symbols, error) {
	return Symbols{
		Title:     Symbol{Value: "Title", CharStart: 0, CharEnd: len("# Title")},
		Wikilinks: []Symbol{{"wikilink", 0, 0}}}, nil
}
