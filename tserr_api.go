package tserr

// All exported error functions are found here, with the exception of NilPtr(), which exists
// in a separate source file. If the function has one argument it is directly provided as
// function argument. If the function has more than one argument, then the arguments are
// provided as a struct, e.g.,
//
//     err := tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: "test1", Y: "test2"})
//
// All error functions first check, if the pointer to the argument struct is nil. If it is
// nil, the error function returns NilPtr(), e.g., 
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

type CheckArgs struct {
	F   string
	Err error
}

func Check(a *CheckArgs) error {
	if a == nil {
		return NilPtr()
	}
	if a.Err == nil {
		return nil
	}
	return errorf(&errmsgCheck, a.F, a.Err)
}

type OpArgs struct {
	Op  string
	Fn  string
	Err error
}

func Op(a *OpArgs) error {
	if a == nil {
		return NilPtr()
	}
	if a.Err == nil {
		return nil
	}
	return errorf(&errmsgOp, a.Op, a.Fn, a.Err)
}

type TypeNotMatchingArgs struct {
	Act  string
	Want string
}

func TypeNotMatching(a *TypeNotMatchingArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgTypeNotMatching, a.Act, a.Want)
}

func NilFailed(op string) error {
	return errorf(&errmsgNilFailed, op)
}

func NotExistent(f string) error {
	return errorf(&errmsgNotExistent, f)
}

func Empty(f string) error {
	return errorf(&errmsgEmpty, f)
}

type NotEqualStrArgs struct {
	X string
	Y string
}

func NotEqualStr(a *NotEqualStrArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotEqualStr, a.X, a.Y)
}
