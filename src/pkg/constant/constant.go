package constant

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

const (
	ErrorDatabaseProblem   = "error database problem"
	ErrorInvalidHttpMethod = "invalid http method"
	ErrorBadRequest        = "error bad request"
	ErrorDataNotFound      = "error data not found"
	ErrorStockNotAvailable = "stock is not available"
)
