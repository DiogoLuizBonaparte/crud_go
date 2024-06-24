package rest_err

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewResErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Resquest",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Resquest",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternarlServerError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal_Server_error",
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not_Found",
		Code:    http.StatusNotFound,
	}
}

func NewForBidenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbiden",
		Code:    http.StatusForbidden,
	}
}
