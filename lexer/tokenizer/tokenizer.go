package tokenizer

import (
	"strconv"
	"strings"
	"unicode"
)

type tokenType int

const (
	Unknown tokenType = iota
	String            // quoted string
	Number            // interger or float
	Word              // word
	Symbol            // symbols
	EOL               // end of line
)

type Token struct {
	SVal string
	NVal float64
	Typ  tokenType
}

func (t *Token) tryNumber() (err error) {
	if t.Typ != Unknown {
		return nil
	}
	t.NVal, err = strconv.ParseFloat(t.SVal, 64)
	if err == nil {
		t.Typ = Number
	} else {
		t.Typ = Word
	}
	return err
}

type Tokenizer interface {
	Next() bool
	Token() *Token
}

func New(str string) Tokenizer {
	return &tokenizer{
		str: strings.TrimSpace(str),
		symbols: map[byte]struct{}{
			',': {},
			'=': {},
			'<': {},
			'>': {},
		},
	}
}

type delimType int

const (
	eol delimType = iota
	eos           // end of search
	symbols
	spaces
	quotes
)

type tokenizer struct {
	str     string
	pos     int    // cursor
	token   *Token // current cached token
	symbols map[byte]struct{}
}

func (t *tokenizer) isSymbol(b byte) bool {
	_, ok := t.symbols[b]
	return ok
}

func (t *tokenizer) nextDelim(pos int) (int, delimType) {
	if pos >= len(t.str) {
		return -1, eol
	}
	for pos < len(t.str) {
		char := t.str[pos]
		switch {
		case t.isSymbol(char):
			return pos, symbols
		case char == '\'':
			return pos, quotes
		case unicode.IsSpace(rune(char)):
			return pos, spaces
		}
		pos++
	}
	return pos, eol
}

func (t *tokenizer) nextQuote(pos int) (int, delimType) {
	if pos >= len(t.str) {
		return len(t.str), eos
	}

	for ; pos < len(t.str); pos++ {
		if t.str[pos] == '\'' {
			return pos, quotes
		}
	}
	return pos, eos
}

func (t *tokenizer) nextNotSpace(pos int) int {
	if pos >= len(t.str) {
		return len(t.str)
	}
	for unicode.IsSpace(rune(t.str[pos])) {
		pos++
	}
	return pos
}

func (t *tokenizer) Token() *Token {
	t.token.tryNumber()
	return t.token
}

func (t *tokenizer) Next() bool {
	delimPos, delimType := t.nextDelim(t.pos)
	nextPos := delimPos
	switch {
	case delimPos == -1 && delimType == eol:
		// can not move further more
		t.token = &Token{Typ: EOL}
		return false
	case delimType == eol:
		t.token = &Token{
			SVal: t.str[t.pos:delimPos],
			Typ:  Unknown, // be determined while Token() calls
		}
	case delimType == spaces:
		if t.pos == delimPos {
			t.pos++
			return t.Next()
		}
		t.token = &Token{
			SVal: t.str[t.pos:delimPos],
			Typ:  Unknown, // be determined while Token() calls
		}
		nextPos = t.nextNotSpace(delimPos + 1)
	case delimType == symbols:
		tpe := Word
		if t.pos == delimPos {
			// catch single symbol, move one character
			nextPos++
			tpe = Symbol
		}
		t.token = &Token{
			SVal: t.str[t.pos:nextPos],
			Typ:  tpe,
		}
	case delimType == quotes:
		nextPos, delimType = t.nextQuote(delimPos + 1)
		if delimType == quotes {
			nextPos++
		}
		t.token = &Token{
			SVal: t.str[t.pos:nextPos],
			Typ:  String,
		}
	}
	t.pos = nextPos
	return true
}
