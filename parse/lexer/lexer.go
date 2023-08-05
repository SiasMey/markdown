package lexer

import (
	"bufio"
	"io"
)

type Token int
const (
	ILLEGAL Token = -1
	EOF Token = 0

	HASH Token = 2
	WIKIOPEN Token = 3
	TAGOPEN Token = 4
	RIGHTBRC Token = 5
)

type Scanner struct {
	r *bufio.Reader
}

var eof = rune(0);

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof;
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if ch == '[' {
		next := s.read()
		if next == '[' {
			return WIKIOPEN, "[["
		} else {
			s.unread()
			return RIGHTBRC, "["
		}
	}

	if ch == '#' {
		next := s.read()
		if next == '[' {
			next2 := s.read()
			if next2 == '[' {
				return TAGOPEN, "#[["
			}
		}
	}

	switch ch {
		case '#':
			return HASH, string(ch)
	}

	return ILLEGAL, string(ch)

}
