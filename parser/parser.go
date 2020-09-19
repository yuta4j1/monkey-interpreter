package parser

import (
	"github.com/yuta4j1/monkey-interpreter/token"
	"github.com/yuta4j1/monkey-interpreter/lexer"
	"github.com/yuta4j1/monkey-interpreter/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken toke.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}