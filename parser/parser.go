package parser

import (
    "fmt"

	"github.com/OPC-16/Monkey/ast"
	"github.com/OPC-16/Monkey/lexer"
	"github.com/OPC-16/Monkey/token"
)

type Parser struct {
    l         *lexer.Lexer

    curToken  token.Token
    peekToken token.Token
    errors    []string
}

// returns an instance of Parser with supplied lexer
func New(l *lexer.Lexer) *Parser {
    p := &Parser{
        l: l,
        errors: []string{},
    }

    // read two tokens, so curToken and peekToken are both set
    p.nextToken()
    p.nextToken()

    return p
}

// helper method that advances both curToken and peekToken
func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

// returns the slice of errors
func (p *Parser) Errors() []string {
    return p.errors
}

// adds new error in errors slice
func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for !p.curTokenIs(token.EOF) {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }

    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
        case token.LET:
            return p.parseLetStatement()
        default:
            return nil
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken}

    if !p.expectPeek(token.IDENT) {
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeek(token.ASSIGN) {
        return nil
    }

    //TODO: we're skipping the expressions until we encounter a semicolon
    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
    return t == p.curToken.Type
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return t == p.peekToken.Type
}

func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.peekError(t)
        return false
    }
}
