package token

// String is a token string
type String string

// Token is a token
type Token struct {
	Type    String
	Literal string
}
