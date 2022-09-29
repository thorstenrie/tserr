package tserr

func ErrChk(f string, err error) error {
	if err == nil {
		return nil
	}
	return errorf(chkFailed, f, err)
}

type ErrOpArgs struct {
	Op string
	Fn string
}

func ErrOp(a *ErrOpArgs, err error) error {
	if err == nil {
		return nil
	}
	if a == nil {
		return ErrNilPtr()
	}
	return errorf(opFailed, a.Op, a.Fn, err)
}

func ErrNilFailed(op string) error {
	return errorf(errNilFailed, op)
}

func ErrNilPtr() error {
	return errorf(nilPtr)
}

func ErrNotExistent(f string) error {
	return errorf(errNilFailed, f)
}
