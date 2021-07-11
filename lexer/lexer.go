package lexer

import (
	"errors"
	"naivedb/lexer/tokenizer"
	"strings"
)

var (
	ErrBadSyntax = errors.New("bad syntax")
)

type empty struct{}

type keywordSet map[string]empty

var keywords keywordSet

func initKeywords() {
	if keywords != nil {
		return
	}
	keywords = keywordSet{}

	sqlKeywords := []string{
		"select", "from", "where", "and", "insert", "into", "values", "delete",
		"drop", "update", "set", "crEate", "table", "int", "double", "varchar",
		"view", "as", "index", "on", "long", "order", "by", "asc", "desc", "sum",
		"count", "avg", "min", "max", "distinct", "group", "add", "sub", "mul",
		"div", "explain", "using", "hash", "btree"}
	for _, key := range sqlKeywords {
		keywords[key] = empty{}
	}
}

type Lexer interface {
	MatchDelim(delimiter string) bool
	MatchNumericConstant() bool
	MatchStringConstant() bool
	MatchKeyword(keyword string) bool
	MatchId() bool

	EatDelim(delimiter string) error
	EatNumericConstant() (float64, error)
	EatStringConstant() (string, error)
	EatKeyword(keyword string) error
	EatId() (string, error)
}

func New(sql string) Lexer {
	l := &lexer{
		tokenizer: tokenizer.New(sql),
	}
	l.tokenizer.Next()
	initKeywords()
	return l
}

type lexer struct {
	tokenizer tokenizer.Tokenizer
}

func (l *lexer) MatchDelim(delimiter string) bool {
	return delimiter == l.tokenizer.Token().SVal
}

func (l *lexer) MatchNumericConstant() bool {
	return l.tokenizer.Token().Typ == tokenizer.Number
}

func (l *lexer) MatchStringConstant() bool {
	return l.tokenizer.Token().Typ == tokenizer.String
}

func (l *lexer) MatchKeyword(keyword string) bool {
	keyword = strings.ToLower(keyword)
	_, ok := keywords[keyword]
	return ok && l.tokenizer.Token().Typ == tokenizer.Word && keyword == strings.ToLower(l.tokenizer.Token().SVal)
}

func (l *lexer) MatchId() bool {
	_, ok := keywords[strings.ToLower(l.tokenizer.Token().SVal)]
	return !ok && l.tokenizer.Token().Typ == tokenizer.Word
}

func (l *lexer) EatDelim(delimiter string) error {
	if !l.MatchDelim(delimiter) {
		return ErrBadSyntax
	}
	l.tokenizer.Next()
	return nil
}

func (l *lexer) EatNumericConstant() (float64, error) {
	if !l.MatchNumericConstant() {
		return 0, ErrBadSyntax
	}
	ret := l.tokenizer.Token().NVal
	l.tokenizer.Next()
	return ret, nil
}

func (l *lexer) EatStringConstant() (string, error) {
	if !l.MatchStringConstant() {
		return "", ErrBadSyntax
	}
	ret := l.tokenizer.Token().SVal
	l.tokenizer.Next()
	return ret, nil
}

func (l *lexer) EatKeyword(keyword string) error {
	if !l.MatchKeyword(keyword) {
		return ErrBadSyntax
	}
	l.tokenizer.Next()
	return nil
}

func (l *lexer) EatId() (string, error) {
	if !l.MatchId() {
		return "", ErrBadSyntax
	}
	ret := l.tokenizer.Token().SVal
	l.tokenizer.Next()
	return ret, nil
}
