package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "if": IF,
    "else": ELSE,
    "true": TRUE,
    "false": FALSE,
    "return": RETURN,
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    //Identifiers + literals
    IDENT = "IDENT"  // add, foobar, x, y, ...
    INT   = "INT"     //3, 234, integers

    // Operators
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"
    LT       = "<"
    GT       = ">"

    // Delimiters
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    LBRACKET = "["
    RBRACKET = "]"

    EQ     = "=="
    NOT_EQ = "!="

    // Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
    IF       = "if"
    ELSE     = "else"
    TRUE     = "true"
    FALSE    = "false"
    RETURN   = "return"

    // String
    STRING   = "STRING"
)

// checks the keywords table to see whether the given identifier is in fact a keyword.
// If it is, it returns the keyword’s TokenType constant.
// If it isn’t, we just get back token.IDENT, which is the TokenType for all user-defined identifiers.
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
