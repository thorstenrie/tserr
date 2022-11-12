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

// CheckArgs holds the required arguments for the error function Check
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

// OpArgs holds the required arguments for the error function Op
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

// TypeNotMatchingArgs holds the required arguments for the error function TypeNotMatching
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

// NotEqualStrArgs holds the required arguments for the error function NotEqualStr
type NotEqualStrArgs struct {
	// X is the string not matching Y
	X string
	// Y is the string not matching X
	Y string
}

// NotEqualStr can be used if two strings are not equal, but expected to be equal.
// A default use case are Test functions.
func NotEqualStr(a *NotEqualStrArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotEqualStr, a.X, a.Y)
}

// ReturnArgs holds the required arguments for the error function Return
type ReturnArgs struct {
	// Op is the operation
	Op string
	// Actual is the actual return value returned by Op
	Actual string
	// Want is the expected return value from Op
	Want string
}

// Return can be used if an operation returns an actual value, but another return value
// is expected. A default use case are Test functions.
func Return(a *ReturnArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgReturn, a.Op, a.Actual, a.Want)
}

// Forbidden can be used if an operation on an object is forbidden
// Argument f is the name of the forbidden object, e.g., directory or filename
func Forbidden(f string) error {
	return errorf(&errmsgForbidden, f)
}

// HigherArgs holds the required arguments for the error function Higher
type HigherArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// LowerBound is the lower bound. Actual is expected to be equal or higher than LowerBound.
	LowerBound int64
}

// Higher can be used if an integer fails to at least be equal or be higher than a defined
// lower bound.
func Higher(a *HigherArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgHigher, a.Var, a.Actual, a.LowerBound)
}

// EqualArgs holds the required arguments for the error function Higher
type EqualArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// Want is the expected value of Var.
	Want int64
}

// Equal can be used if an integer fails to be equal to an expected value.
func Equal(a *EqualArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqual, a.Var, a.Actual, a.Want)
}
