package cerr

import "net/http"

type Code string

func (c Code) String() string { return string(c) }

func (c Code) HTTPCode() int {
	httpCode, exist := codes[c]
	if !exist {
		return http.StatusInternalServerError
	}

	return httpCode
}

var (
	CodeInternal           Code = "ERR_INTERNAL"
	CodeNotPermitted       Code = "ERR_NOT_PERMITTED"
	CodeValidate           Code = "ERR_VALIDATE"
	CodeNotFound           Code = "ERR_NOT_FOUND"
	CodeInitInProgress     Code = "ERR_INIT_IN_PROGRES"
	CodeInvalidCredentials Code = "ERR_INVALID_CREDENTIALS"
)

var codes = map[Code]int{
	CodeInternal:           http.StatusInternalServerError,
	CodeValidate:           http.StatusBadRequest,
	CodeInitInProgress:     http.StatusBadRequest,
	CodeInvalidCredentials: http.StatusBadRequest,
	CodeNotPermitted:       http.StatusForbidden,
}
