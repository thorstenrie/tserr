package tserr

var (
	nilPtr = errmsg{0, 500, "nil pointer"}
)

func NilPtr() error {
	return errorf(&nilPtr)
}
