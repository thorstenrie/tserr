package tserr

// All exported error functions are implemented here, with the exception of NilPtr, which exists
// in a separate source file. If the function has one argument it is directly provided as
// function argument. If the function has more than one argument, then the arguments are
// provided as a struct, e.g.,
//
//     err := tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: "test1", Y: "test2"})
//
// All error functions first check, if the pointer to the argument struct is nil. If it is
// nil, the error function returns NilPtr, e.g.,
//
//     if a == nil {
//	       return NilPtr()
//     }
//
// If the argument struct contains a field of type error, is is checked next. If the error field is
// nil, then the error function returns nil, e.g.,
//
//     if a.Err == nil {
//         return nil
//     }
//
// Otherwise, it returns the corresponding error message, e.g.,
//
//     return errorf(&errmsgNotEqualStr, a.X, a.Y)

// CheckArgs holds the required arguments for error function Check
type CheckArgs struct {
	// F is the name of the object causing the failed check, e.g., a filename
	F string
	// Err is the error causing the failed check, e.g., file is a directory
	Err error
}

// Check can be used for negative validations on an object
func Check(a *CheckArgs) error {
	if a == nil {
		return NilPtr()
	}
	if a.Err == nil {
		return nil
	}
	return errorf(&errmsgCheck, a.F, a.Err)
}

// OpArgs holds the required arguments for error function Op
type OpArgs struct {
	// Op is the name of the failed operation, e.g., WriteStr
	Op string
	// Fn is the name of the object passed to the operation, e.g., a filename
	Fn string
	// Err is the error retrieved from the failed operation, e.g., file does not exist
	Err error
}

// Op can be used for failed operations on an object
func Op(a *OpArgs) error {
	if a == nil {
		return NilPtr()
	}
	if a.Err == nil {
		return nil
	}
	return errorf(&errmsgOp, a.Op, a.Fn, a.Err)
}

// TypeNotMatchingArgs holds the required arguments for error function TypeNotMatching
type TypeNotMatchingArgs struct {
	// Act is the name of the actual type of the object, e.g., file
	Act string
	// Want is the name of the expected, wanted or required type the object should be, e.g., directory
	Want string
}

// TypeNotMatching can be used if the type of an object does not match the expected type
func TypeNotMatching(a *TypeNotMatchingArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgTypeNotMatching, a.Act, a.Want)
}

// NilFailed can be used if the function implementing an operation returns nil, but an error is expected.
// A default use case are Test functions.
// Argument op is the name of the operation, e.g., ExistsFile
func NilFailed(op string) error {
	return errorf(&errmsgNilFailed, op)
}

// NotExistent can be used if an required object does not exist, e.g., a file.
// Argument f is the name of the object, e.g., filename
func NotExistent(f string) error {
	return errorf(&errmsgNotExistent, f)
}

// Empty can be used if an required object is empty but not allowed to be empty, e.g., an input argument of type string.
// Argument f is the name of the empty object, e.g., filename
func Empty(f string) error {
	return errorf(&errmsgEmpty, f)
}

// NotEqualStrArgs holds the required arguments for error function NotEqualStr
type NotEqualStrArgs struct {
	// X is the string not matching Y
	X string
	// Y is the string not matching X
	Y string
}

// NotEqualStr can be used if to strings are not equal, but expected to be equal.
// A default use case are Test functions.
func NotEqualStr(a *NotEqualStrArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotEqualStr, a.X, a.Y)
}
