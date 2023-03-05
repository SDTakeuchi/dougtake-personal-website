package constants

type ResponseMessage string

var (
	AuthenticationFailed ResponseMessage = "authentication failed"
	InvalidToken         ResponseMessage = "token is invalid"
	ExpiredToken         ResponseMessage = "token is already expired"

	PasswordTooShort  ResponseMessage = "password must be at least 8 characters"
	IncorrectPassword ResponseMessage = "incorrect password"

	DefaultErrorMessage ResponseMessage = "unexpected error has occurred: contact to the owner"
	RecordNotFound      ResponseMessage = "record not found"
	FailedToBindQuery   ResponseMessage = "failed to bind query: wrong format"
	InvalidParams       ResponseMessage = "invalid parameters"
)

func (r ResponseMessage) String() string {
	return string(r)
}
