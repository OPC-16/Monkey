package ast

import (
	"bytes"

	"github.com/OPC-16/Monkey/token"
)

// Every node in our AST has to implement the Node interface, meaning it has to provide a TokenLiteral() method that returns the literal value of
// the token it’s associated with.
type Node interface {
    TokenLiteral() string
    String()       string   //prints AST nodes for debugging and to compare them with other AST nodes
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

// This Program node is going to be the root node of every AST our parser produces.
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

// it creates a buffer and writes the return value of each Statement's String() method to it. And then it returns the buffer as a string.
// With help of String() methods, we can just call String() on *ast.Program and get our whole program back as a string
func (p *Program) String() string {
    var out bytes.Buffer

    for _, s := range p.Statements {
        out.WriteString(s.String())
    }

    return out.String()
}

type LetStatement struct {
    Token token.Token   // the token.LET token
    Name  *Identifier
    Value Expression
}

func (ls *LetStatement) statementNode() { }
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
    var out bytes.Buffer

    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")

    if ls.Value != nil {
        out.WriteString(ls.Value.String())
    }

    out.WriteString(";")
    return out.String()
}

type Identifier struct {
    Token token.Token   // the token.IDENT token
    Value string
}

func (i *Identifier) expressionNode() { }
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

type ReturnStatement struct {
    Token       token.Token   // the 'Return' token
    ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() { }
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
    var out bytes.Buffer
    
    out.WriteString(rs.TokenLiteral() + " ")

    if rs.ReturnValue != nil {
        out.WriteString(rs.ReturnValue.String())
    }

    out.WriteString(";")
    return out.String()
}

// it fulfills the ast.Statement interface, which means we can add it to the Statements slice of ast.Program
type ExpressionStatement struct {
    Token      token.Token   // the first token of the expression
    Expression Expression
}

func (es *ExpressionStatement) statementNode() { }
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}

type IntegerLiteral struct {
    Token token.Token
    Value int64
}

func (il *IntegerLiteral) expressionNode() { }
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

type PrefixExpression struct {
    Token    token.Token   // the prefix token, e.g. ! or -
    Operator string
    Right    Expression
}

func (pe *PrefixExpression) expressionNode() { }
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
    var out bytes.Buffer

    out.WriteString("(")
    out.WriteString(pe.Operator)
    out.WriteString(pe.Right.String())
    out.WriteString(")")

    return out.String()
}
