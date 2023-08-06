package lexer

import (
	"bufio"
	"io"
	"log"
)

type Token int

const (
	ILLEGAL Token = -1
	EOF     Token = 0

	HASH      Token = 2
	WIKIOPEN  Token = 3
	LEFTBRC   Token = 4
	RIGHTBRC  Token = 5
	WIKICLOSE Token = 6
)

type Scanner struct {
	r *bufio.Reader
}

var eof = rune(0)
var logger = log.Default()

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	err := s.r.UnreadRune()
	if err != nil {
		logger.Print("Error during unread rune", err)
	}
}

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if ch == '[' {
		next := s.read()
		if next == '[' {
			return WIKIOPEN, "[["
		} else {
			s.unread()
			return LEFTBRC, "["
		}
	}

	if ch == ']' {
		next := s.read()
		if next == ']' {
			return WIKICLOSE, "]]"
		} else {
			s.unread()
			return RIGHTBRC, string(ch)
		}
	}

	switch ch {
	case '#':
		return HASH, string(ch)
	}

	return ILLEGAL, string(ch)

}
