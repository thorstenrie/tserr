// The source file holds the NilPtr error function. Since it is self-used internally
// in the package, it is provided in a separate source file.
package tserr

// nilPtr error message is id 0 with error code 500 and a simple
// error message without verbs.
var (
	nilPtr = errmsg{0, 500, "nil pointer"}
)

// NilPtr just provides the error message and does not have arguments.
func NilPtr() error {
	return errorf(&nilPtr)
}
