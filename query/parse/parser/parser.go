package parser

import (
	"naivedb/query/parse/lexer"
)

type Parser interface {
	QueryCommand() *QueryData
	ExecCommand() *ExecData
}

func New(sql string) Parser {
	return &parser{lexer.New(sql)}
}

type parser struct {
	lexer lexer.Lexer
}

func (p *parser) QueryCommand() *QueryData {
	return nil
}

func (p *parser) ExecCommand() *ExecData {
	return nil
}
