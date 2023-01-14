package constants

type ResponseMessage string

var (
	DefaultErrorMessage ResponseMessage = "unexpected error has occurred: contact to the owner"
)

func (r ResponseMessage) String() string {
	return string(r)
}
