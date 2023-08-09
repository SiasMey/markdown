package lexer

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

type TokenType int

const (
	ILLEGAL TokenType = -1
	EOF     TokenType = 0

	HASH     TokenType = 2
	LEFTBRK  TokenType = 4
	RIGHTBRK TokenType = 5
	LEFTPRN  TokenType = 6
	RIGHTPRN TokenType = 11
	TEXT     TokenType = 7
	WS       TokenType = 8
	NL       TokenType = 9
	TICK     TokenType = 10
)

type Token struct {
	TypeType TokenType
	Lit string
	LineNr int
	Length int
}

type Scanner struct {
	r *bufio.Reader
}

var eof = rune(0)
var logger = log.Default()

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() Token {
	token, lit := s.scanNext()

	return Token{
		TypeType: token,
		Lit: lit,
		LineNr: 0,
		Length: 0,
	}
}

func (s *Scanner) scanNext() (TokenType, string) {
	ch := s.read()

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
	case '`':
		return TICK, string(ch)
	case '[':
		return LEFTBRK, string(ch)
	case ']':
		return RIGHTBRK, string(ch)
	case '(':
		return LEFTPRN, string(ch)
	case ')':
		return RIGHTPRN, string(ch)
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

func (s *Scanner) scanNewLine() (TokenType, string) {
	nl := s.read()

	if nl == '\r' {
		win := s.read()
		if win == '\n' {
			return NL, string(nl) + string(win)
		} else {
			s.unread()
			return NL, string(nl)
		}
	}

	return NL, string(nl)
}

func (s *Scanner) scanHash() (TokenType, string) {
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

func (s *Scanner) scanWhiteSpace() (TokenType, string) {
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

func (s *Scanner) scanText() (TokenType, string) {
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
