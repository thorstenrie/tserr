package tserr

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
	return errorf(chkFailed, a.F, a.Err)
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
	return errorf(opFailed, a.Op, a.Fn, a.Err)
}

type TypeArgs struct {
	Act  string
	Want string
}

func TypeNotMatching(a *TypeArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(typeNotEqual, a.Act, a.Want)
}

func NilFailed(op string) error {
	return errorf(errNilFailed, op)
}

func NilPtr() error {
	return errorf(nilPtr)
}

func NotExistent(f string) error {
	return errorf(errNilFailed, f)
}

func Empty(f string) error {
	return errorf(cannotBeEmpty, f)
}

type TypeNotEqualStr struct {
	X string
	Y string
}

func NotEqualStr(a *TypeNotEqualStr) error {
	return errorf(strNotEqual, a.X, a.Y)
}
