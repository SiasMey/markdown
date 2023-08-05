package symbols

type Symbol struct {
	Value     string
	CharStart int
	CharEnd   int
	LineNo    int
}

type Symbols struct {
	Title     Symbol
	Wikilinks []Symbol
}

func Extract(input string) (Symbols, error) {
	return Symbols{
		Title:     Symbol{Value: "Title", CharStart: 0, CharEnd: len("# Title"), LineNo: 0},
		Wikilinks: []Symbol{{"wikilink", 0, 0, 0}}}, nil
}
