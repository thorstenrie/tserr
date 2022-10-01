package tserr

import (
	"fmt"
	"testing"
)

var (
	strFoo string = "tserr_foo"
	errFoo error  = fmt.Errorf(strFoo)
)

func TestCheckNil(t *testing.T) {
	if err := Check(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestCheckNilErr(t *testing.T) {
	a := CheckArgs{Err: nil}
	if err := Check(&a); err != nil {
		t.Errorf(errNNil)
	}
}

func TestCheck(t *testing.T) {
	a := CheckArgs{
		F:   strFoo,
		Err: errFoo,
	}
	em := &errmsgCheck
	err := Check(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a.F,
			a.Err,
		)),
	}
	testEqualJson(t, err, &emsg)
}

func TestOpNil(t *testing.T) {
	if err := Op(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestOpNilErr(t *testing.T) {
	a := OpArgs{Err: nil}
	if err := Op(&a); err != nil {
		t.Errorf(errNNil)
	}
}

func TestOp(t *testing.T) {
	a := OpArgs{
		Op:  strFoo,
		Fn:  strFoo,
		Err: errFoo,
	}
	em := &errmsgOp
	err := Op(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a.Op,
			a.Fn,
			a.Err,
		)),
	}
	testEqualJson(t, err, &emsg)
}

func TestTypeNotMatchingNil(t *testing.T) {
	if err := TypeNotMatching(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestTypeNotMatching(t *testing.T) {
	a := TypeNotMatchingArgs{
		Act:  strFoo,
		Want: strFoo,
	}
	em := &errmsgTypeNotMatching
	err := TypeNotMatching(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a.Act,
			a.Want,
		)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotEqualStrNil(t *testing.T) {
	if err := NotEqualStr(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestNotEqualStr(t *testing.T) {
	a := NotEqualStrArgs{
		X: strFoo,
		Y: strFoo + strFoo,
	}
	em := &errmsgNotEqualStr
	err := NotEqualStr(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a.X,
			a.Y,
		)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNilFailed(t *testing.T) {
	a := strFoo
	em := &errmsgNilFailed
	err := NilFailed(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a,
		)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotExistent(t *testing.T) {
	a := strFoo
	em := &errmsgNotExistent
	err := NotExistent(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a,
		)),
	}
	testEqualJson(t, err, &emsg)
}

func TestEmpty(t *testing.T) {
	a := strFoo
	em := &errmsgEmpty
	err := Empty(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a,
		)),
	}
	testEqualJson(t, err, &emsg)
}
