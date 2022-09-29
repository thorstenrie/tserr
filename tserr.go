package tserr

func Check(f string, err error) error {
	if err == nil {
		return nil
	}
	return errorf(chkFailed, f, err)
}

type ErrOpArgs struct {
	Op string
	Fn string
}

func Op(a *ErrOpArgs, err error) error {
	if err == nil {
		return nil
	}
	if a == nil {
		return NilPtr()
	}
	return errorf(opFailed, a.Op, a.Fn, err)
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

func NotEqualStr(a string, b string) error {
	return errorf(strNotEqual, a, b)
}
