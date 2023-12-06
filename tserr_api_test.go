package tserr

// All tests for all exported error functions are implemented here, with the exception of NilPtr(). The
// test of NilPtr() exists in a separate source file. Each error function runs through up to three
// tests.
//
//     1) Test for all functions: returned error is not nil, holds an error message in valid
//        JSON format and equals the expected error message
//     2) Additional test for error functions with multiple arguments passed in a struct:
//        Check for returned error if pointer to argument struct is nil
//     3) Additional test for error functions with multiple arguments passed in a struct and
//        one argument is of type error: Check if return value is nil in case provided
//        error in argument struct is nil.
//
// The structure of all test functions follows the same pattern. For an example, please see
// tests for Check: TestCheck, TestCheckNil, TestCheckNilErr
//
// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.

// Import standard library packages
import (
	"fmt"     // fmt
	"testing" // testing
)

// testcases for types string and error
var (
	strFoo   string  = "tserr_foo"        // testcase type string
	errFoo   error   = fmt.Errorf(strFoo) // testcase type error
	intFoo   int64   = 1234               // testcase type int64
	floatFoo float64 = 1234               // testcase type float64
)

func TestCheckNil(t *testing.T) {
	if err := Check(nil); err == nil {
		t.Errorf(errNil)
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
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.F, a.Err)),
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
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestOpNil(t *testing.T) {
	if err := Op(nil); err == nil {
		t.Errorf(errNil)
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
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Op, a.Fn, a.Err)),
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
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotNil(t *testing.T) {
	a := strFoo
	em := &errmsgNotNil
	err := NotNil(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
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
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotEmpty(t *testing.T) {
	a := strFoo
	em := &errmsgNotEmpty
	err := NotEmpty(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestEqualStrNil(t *testing.T) {
	if err := EqualStr(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestEqualStr(t *testing.T) {
	a := EqualStrArgs{
		Var:    strFoo,
		Actual: strFoo,
		Want:   strFoo,
	}
	em := &errmsgEqualStr
	err := EqualStr(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Var, a.Actual, a.Want)),
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
		Actual: strFoo,
		Want:   strFoo,
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
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Actual, a.Want)),
	}
	testEqualJson(t, err, &emsg)
}

func TestForbidden(t *testing.T) {
	a := strFoo
	em := &errmsgForbidden
	err := Forbidden(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestReturnNil(t *testing.T) {
	if err := Return(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestReturn(t *testing.T) {
	a := ReturnArgs{
		Op:     strFoo,
		Actual: strFoo,
		Want:   strFoo,
	}
	em := &errmsgReturn
	err := Return(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Op, a.Actual, a.Want)),
	}
	testEqualJson(t, err, &emsg)
}

func TestHigherNil(t *testing.T) {
	if err := Higher(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestHigher(t *testing.T) {
	a := HigherArgs{
		Var:        strFoo,
		Actual:     intFoo,
		LowerBound: intFoo,
	}
	em := &errmsgHigher
	err := Higher(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Var, a.Actual, a.LowerBound)),
	}
	testEqualJson(t, err, &emsg)
}

func TestEqualNil(t *testing.T) {
	if err := Equal(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestEqual(t *testing.T) {
	a := EqualArgs{
		Var:    strFoo,
		Actual: intFoo,
		Want:   intFoo,
	}
	em := &errmsgEqual
	err := Equal(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Var, a.Actual, a.Want)),
	}
	testEqualJson(t, err, &emsg)
}

func TestLowerNil(t *testing.T) {
	if err := Lower(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestLower(t *testing.T) {
	a := LowerArgs{
		Var:    strFoo,
		Actual: intFoo,
		Want:   intFoo,
	}
	em := &errmsgLower
	err := Lower(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Var, a.Actual, a.Want)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotSet(t *testing.T) {
	a := strFoo
	em := &errmsgNotSet
	err := NotSet(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotAvailableNil(t *testing.T) {
	if err := NotAvailable(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestNotAvailable(t *testing.T) {
	a := NotAvailableArgs{
		S:   strFoo,
		Err: errFoo,
	}
	em := &errmsgNotAvailable
	err := NotAvailable(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.S, a.Err)),
	}
	testEqualJson(t, err, &emsg)
}

func TestEqualfNil(t *testing.T) {
	if err := Equalf(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestEqualf(t *testing.T) {
	a := EqualfArgs{
		Var:    strFoo,
		Actual: floatFoo,
		Want:   floatFoo,
	}
	em := &errmsgEqualf
	err := Equalf(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.Var, a.Actual, a.Want)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNonPrintable(t *testing.T) {
	a := strFoo
	em := &errmsgNonPrintable
	err := NonPrintable(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestNotEqualNil(t *testing.T) {
	if err := NotEqual(nil); err == nil {
		t.Errorf(errNil)
	}
}

func TestNotEqual(t *testing.T) {
	a := NotEqualArgs{
		X: strFoo,
		Y: strFoo,
	}
	em := &errmsgNotEqual
	err := NotEqual(&a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a.X, a.Y)),
	}
	testEqualJson(t, err, &emsg)
}

func TestDuplicate(t *testing.T) {
	a := strFoo
	em := &errmsgDuplicate
	err := Duplicate(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}

func TestLocked(t *testing.T) {
	a := strFoo
	em := &errmsgLocked
	err := Locked(a)
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(em.M, a)),
	}
	testEqualJson(t, err, &emsg)
}
