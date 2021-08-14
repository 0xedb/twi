package token

// TokenString is a token string
type TokenString string

// Token is a token
type Token struct {
	Type    TokenString
	Literal string
}
