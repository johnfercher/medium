package apierror

type ApiError interface {
	Name() string
	Code() int
}

type apiError struct {
	name string
	code int
}

func New(name string, code int) ApiError {
	return &apiError{name, code}
}

func (a *apiError) Name() string {
	return a.name
}

func (a *apiError) Code() int {
	return a.code
}
