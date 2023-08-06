package lexer

import (
	"bufio"
	"bytes"
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
	TEXT      Token = 7
	WS        Token = 8
	NL Token = 9
)

type Scanner struct {
	r *bufio.Reader
}

var eof = rune(0)
var logger = log.Default()

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (Token, string) {
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

	if isText(ch) {
		s.unread()
		return s.scanText()
	}
	
	if isWhiteSpace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	}

	if isNewLine(ch) {
		s.unread()
		return s.scanNewLine()
	}

	switch ch {
	case '#':
		s.unread()
		return s.scanHash()
	case eof:
		return EOF, string(ch)
	}

	return ILLEGAL, string(ch)

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
		logger.Print("ERROR: unread rune", err)
	}
}

func isNewLine(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func isPunc(ch rune) bool {
	return ch == '.' || ch == ',' || ch == '-' || ch == '_'
}

func isNum(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isText(ch rune) bool {
	return isAlpha(ch) || isNum(ch) || isPunc(ch)
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func (s *Scanner) scanNewLine() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNewLine(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return NL, buf.String()
}

func (s *Scanner) scanHash() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch != '#' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return HASH, buf.String()
}

func (s *Scanner) scanWhiteSpace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhiteSpace(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanText() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isText(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return TEXT, buf.String()
}
