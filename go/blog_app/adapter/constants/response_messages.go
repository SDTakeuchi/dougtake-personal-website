package constants

type ResponseMessage string

var (
	DefaultErrorMessage ResponseMessage = "unexpected error has occurred: contact to the owner"
	RecordNotFound      ResponseMessage = "record not found"
	FailedToBindQuery   ResponseMessage = "failed to bind query: wrong format"
)

func (r ResponseMessage) String() string {
	return string(r)
}
