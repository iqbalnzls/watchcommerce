package constant

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

const (
	ErrorDatabaseProblem   = "database problem"
	ErrorInvalidHttpMethod = "invalid http method"
	ErrorBadRequest        = "bad request"
	ErrorDataNotFound      = "data not found"
	ErrorStockNotAvailable = "stock is not available"
	ErrorProductNotFound   = "product not found"
)
