// All exported error functions are implemented here, with the exception of NilPtr, which exists
// in a separate source file. If the function has one argument it is directly provided as
// function argument. If the function has more than one argument, then the arguments are
// provided as a struct, e.g.,
//
//	err := tserr.EqualStr(&tserr.EqualStrArgs{X: "test1", Y: "test2"})
//
// All error functions first check, if the pointer to the argument struct is nil. If it is
// nil, the error function returns NilPtr, e.g.,
//
//	if a == nil {
//	    return NilPtr()
//	}
//
// Otherwise, it returns the corresponding error message, e.g.,
//
//	return errorf(&errmsgEqualStr, a.X, a.Y)
//
// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tserr

// CheckArgs holds the required arguments for the error function Check
type CheckArgs struct {
	// F is the name of the object causing the failed check, for example, a filename
	F string
	// Err is the error causing the failed check, for example, file is a directory
	Err error
}

// Check can be used for negative validations on an object.
func Check(a *CheckArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgCheck, a.F, a.Err)
}

// NotExistent can be used if an required object does not exist, for example, a file.
// F is the name of the object, for example, key name
func NotExistent(F string) error {
	return errorf(&errmsgNotExistent, F)
}

// OpArgs holds the required arguments for the error function Op
type OpArgs struct {
	// Op is the name of the failed operation, for example, WriteStr
	Op string
	// Fn is the name of the object passed to the operation, for example, a filename
	Fn string
	// Err is the error retrieved from the failed operation, for example, file does not exist
	Err error
}

// Op can be used for failed operations on an object.
func Op(a *OpArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgOp, a.Op, a.Fn, a.Err)
}

// NilFailed can be used if the function implementing an operation returns nil, but an error is expected. A default use case are Test functions.
// Op is the name of the operation, for example, ExistsFile
func NilFailed(Op string) error {
	return errorf(&errmsgNilFailed, Op)
}

// Empty can be used if a required object is empty but not allowed to be empty, for example, an input argument of type string.
// F is the name of the empty object, for example, filename
func Empty(F string) error {
	return errorf(&errmsgEmpty, F)
}

// NotEmpty can be used if a required object is not empty but expected to be empty, for example, an output argument of type string.
// F is the name of the object expected to be empty, for example, filename
func NotEmpty(F string) error {
	return errorf(&errmsgNotEmpty, F)
}

// EqualStrArgs holds the required arguments for the error function EqualStr
type EqualStrArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual string
	// Want is the expected value of Var
	Want string
}

// EqualStr can be used if a string fails to be equal to an expected string.
func EqualStr(a *EqualStrArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqualStr, a.Var, a.Actual, a.Want)
}

// TypeNotMatchingArgs holds the required arguments for the error function TypeNotMatching
type TypeNotMatchingArgs struct {
	// Actual is the name of the actual type of the object, for example, a file
	Actual string
	// Want is the name of the expected, wanted or required type the object should be, for example, a directory
	Want string
}

// TypeNotMatching can be used if the type of an object does not match the expected type
func TypeNotMatching(a *TypeNotMatchingArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgTypeNotMatching, a.Actual, a.Want)
}

// Forbidden can be used if an operation on an object is forbidden.
// F is the name of the forbidden object, for example, a directory or filename
func Forbidden(F string) error {
	return errorf(&errmsgForbidden, F)
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

// Return can be used if an operation returns an actual value, but another return value is expected.
func Return(a *ReturnArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgReturn, a.Op, a.Actual, a.Want)
}

// HigherArgs holds the required arguments for the error function Higher
type HigherArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual of Var
	Actual int64
	// LowerBound is the lower bound. Actual is expected to be equal or higher than Lowerbound
	LowerBound int64
}

// Higher can be used if an integer fails to at least be equal or be higher than a defined lower bound.
func Higher(a *HigherArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgHigher, a.Var, a.Actual, a.LowerBound)
}

// EqualArgs holds the required arguments for the error function Equal
type EqualArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// Want is the expected value of Var
	Want int64
}

// Equal can be used if an integer tails to be equal to an expected value.
func Equal(a *EqualArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqual, a.Var, a.Actual, a.Want)
}

// LowerArgs holds the required arguments for the error function Lower
type LowerArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// Want is the expected value of Var
	Want int64
}

// Lower can be used if an integer fails to be lower than a defined higher bound.
func Lower(a *LowerArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgLower, a.Var, a.Actual, a.Want)
}

// NotSet can be used if a required object is not set, for example, an environment variable.
// F is the name of the object, for example, the name of the environment variable
func NotSet(F string) error {
	return errorf(&errmsgNotSet, F)
}

// NotAvailableArgs holds the required arguments for the error function NotAvailable
type NotAvailableArgs struct {
	// S is the name of the service not available
	S string
	// Err is the error provided by the service
	Err error
}

// NotAvailable can be used if a service is not available.
func NotAvailable(a *NotAvailableArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotAvailable, a.S, a.Err)
}

// EqualfArgs holds the required arguments for the error function Equalf
type EqualfArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual float64
	// Want is the expected value of Var
	Want float64
}

// Equalf can be used if a float value is not equal to an expected value
func Equalf(a *EqualfArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqualf, a.Var, a.Actual, a.Want)
}

// NonPrintable can be used if a string is allowed to only contain printable runes, but actually contains non-printable runes.
// F is the name of the string allowed to only contain printable runes
func NonPrintable(F string) error {
	return errorf(&errmsgNonPrintable, F)
}

// NotEqualArgs holds the required arguments for the error function NotEqual
type NotEqualArgs struct {
	// Name of the variable equal to Y
	X string
	// Name of the variable equal to X
	Y string
}

// NotEqual can be used if two variables are equal but not expected to be equal.
func NotEqual(a *NotEqualArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotEqual, a.X, a.Y)
}

// Duplicate can be used if an object already exists, but is only allowed to exist once, for example, a key.
// F is the name of the object which already exists, for example, the name of a key
func Duplicate(F string) error {
	return errorf(&errmsgDuplicate, F)
}

// Locked can be used if a service is locked, for example, because it is still running.
// S is the name of the service which is locked
func Locked(S string) error {
	return errorf(&errmsgLocked, S)
}
