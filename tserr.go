package tserr

import "fmt"

type errmsg struct {
	Id int    `json:"id"`      // id
	C  int    `json:"code"`    // error code
	M  string `json:"message"` // error message
}

type errwrap struct {
	E errmsg `json:"error"`
}

var (
	errformat string = "{" +
		"\"error\":{" +
		"\"id\":%d," +
		"\"code\":%d," +
		"\"message\":\"%w\"" +
		"}" +
		"}"
)

func errorf(e *errmsg, a ...any) error {
	if e == nil {
		return fmt.Errorf(errformat, nilPtr.Id, nilPtr.C, nilPtr.M)
	}
	return fmt.Errorf(errformat, e.Id, e.C, fmt.Errorf(e.M, a...))
}
