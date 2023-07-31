package token

type TokenType string 

type Token struct {
    Type TokenType
    Literal string
}

// Define token types as constants

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers and literals
    IDENT = "IDENT"
    INT = "INT"

    // Operators 
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LEFTPAREN = "("
    RIGHTPAREN = ")"
    LEFTBRACE = "{"
    RIGHTBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"

)
