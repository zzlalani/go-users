package enumeration

import "net/http"

type AppError struct {
	Code		int		`json:"code"`
	Error		error	`json:"error"`
	Name		string	`json:"name"`
	StatusCode	int		`json:"status_code"`
}

func newError(_code int, _error error, _name string, _statusCode int) AppError {
	return AppError{
		Code: _code,
		Error: _error,
		Name: _name,
		StatusCode: _statusCode,
	}
}

func BadRequest(_code int, _error error) AppError {
	return newError(_code, _error, "BadRequest", http.StatusBadRequest)
}

func Conflict(_code int, _error error) AppError {
	return newError(_code, _error, "Conflict", http.StatusConflict)
}

func Forbidden(_code int, _error error) AppError {
	return newError(_code, _error, "Forbidden", http.StatusForbidden)
}

func InternalServerError(_code int, _error error) AppError {
	return newError(_code, _error, "InternalServerError", http.StatusInternalServerError)
}

func NotFound(_code int, _error error) AppError {
	return newError(_code, _error, "NotFound", http.StatusNotFound)
}

func Unauthorized(_code int, _error error) AppError {
	return newError(_code, _error, "Unauthorized", http.StatusUnauthorized)
}

func UnprocessableEntity(_code int, _error error) AppError {
	return newError(_code, _error, "UnprocessableEntity", http.StatusUnprocessableEntity)
}