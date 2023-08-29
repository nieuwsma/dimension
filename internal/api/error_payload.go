package api

import (
	"net/http"
)

type InternalServerError struct {
}

func (e *InternalServerError) Error() string { return "internal server error" }

type InvalidPool struct {
}

func (e *InvalidPool) Error() string { return "invalid pool" }

type StorageFailed struct {
}

func (e *StorageFailed) Error() string { return "internal storage error" }

type BadRequest struct {
}

func (e *BadRequest) Error() string { return "bad request" }

type NotFound struct {
}

func (e *NotFound) Error() string { return "not found" }

type Forbidden struct {
}

func (e *Forbidden) Error() string { return "Forbidden" }

func GetFormattedErrorMessage(err error, code int) Problem7807 {

	var JSONerror Problem7807

	JSONerror.Type_ = "about:blank"
	JSONerror.Instance = ""
	JSONerror.Status = code
	if err != nil {
		JSONerror.Detail = err.Error()
	} else {
		JSONerror.Detail = "unknown error - could not parse from error object"
	}
	JSONerror.Title = http.StatusText(code)

	return JSONerror
}
