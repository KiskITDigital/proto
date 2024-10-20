package cerr

import "errors"

var (
	ErrPermission = errors.New("not permitted")
)
