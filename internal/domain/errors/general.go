package errors

type ErrorCodeMsg struct {
	Code	string
	Message	string
}

var (
	// Token/JWT
	MissingToken = ErrorCodeMsg{Code: "MISSING_TOKEN", Message: "Missing Token."}
	InvalidTokenFormat = ErrorCodeMsg{Code: "INVALID_TOKEN_FORMAT", Message: "Invalid Token Format."}
	InvalidTokenClaims = ErrorCodeMsg{Code: "INVALID_TOKEN_CLAIMS", Message: "Invalid Token Claims."}
	UnexpectedSigningMethod = ErrorCodeMsg{Code: "UNEXPECTED_SIGNING_METHOD", Message: "Unexpected signing method."}
)