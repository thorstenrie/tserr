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
