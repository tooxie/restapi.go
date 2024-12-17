package errors

type errorName string

const (
	_ errorName = "Ok"

	// Server errors
	ErrorRetrievingBooks     = "ErrorRetrievingBooks"
	ErrorSavingBook          = "ErrorSavingBook"
	ErrorUpdatingBook        = "ErrorUpdatingBook"
	InternalServerError      = "InternalServerError"
	JwtTokenGenerationFailed = "JwtTokenGenerationFailed"

	// Client errors
	BookNotFound       = "BookNotFound"
	EmailAlreadyInUse  = "EmailAlreadyInUse"
	Forbidden          = "Forbidden"
	InvalidCredentials = "InvalidCredentials"
	InvalidHostHeader  = "InvalidHostHeader"
	InvalidInput       = "InvalidInput"
	JwtInvalidToken    = "JwtInvalidToken"
	JwtTokenExpired    = "JwtTokenExpired"
	Unauthorized       = "Unauthorized"
)
