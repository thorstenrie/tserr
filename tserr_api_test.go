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

// Import standard library packages
import (
	"fmt"     // fmt
	"testing" // testing
)

// testcases for types string and error
var (
	strFoo string = "tserr_foo"        // testcase type string
	errFoo error  = fmt.Errorf(strFoo) // testcase type error
)

// TestCheckNil tests if an error is returned by Check in case
// the argument pointer is nil. If the returned error is nil
// then the test reports an error.
func TestCheckNil(t *testing.T) {
	// Set argument to nil and test fails if returned error is nil
	if err := Check(nil); err == nil {
		t.Errorf(errNil)
	}
}

// TestCheckNilErr tests if the return value of Check is nil
// in case the argument of type error is nil. If the return
// value is not nil, then the test reports an error.
func TestCheckNilErr(t *testing.T) {
	// Create Check argument struct with Err of type error
	// set to nil
	a := CheckArgs{Err: nil}
	// Return error if Check returns error and not nil.
	if err := Check(&a); err != nil {
		t.Errorf(errNNil)
	}
}

// TestCheck tests, if the returned error of Check() is not nil,
// holds an error message in valid JSON format and if the returned error
// message equals the expected error message
func TestCheck(t *testing.T) {
	// Create Check argument struct a
	a := CheckArgs{
		F:   strFoo,
		Err: errFoo,
	}
	// Get a pointer to the corresponding error message with verbs
	em := &errmsgCheck
	// Retrieve error from Check() for argument struct a
	err := Check(&a)
	// Test fails if returned error is nil
	if err == nil {
		t.Fatal(errNil)
	}
	// Test if returned error message is in a valid JSON format
	testValidJson(t, err)
	// Create emsg of type errmsg with verbs filled by provided arguments
	// from argument struct a. emsg will be transformed in the expected, wanted
	// error message and compared to the actual returned error message of
	// Check()
	emsg := errmsg{
		em.Id,
		em.C,
		fmt.Sprintf("%v", fmt.Errorf(
			em.M,
			a.F,
			a.Err,
		)),
	}
	// Test if returned error message equals expected error message. The
	// expected error message is generated based on emsg.
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
