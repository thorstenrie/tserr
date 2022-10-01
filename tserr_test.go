package tserr

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	errNil  string = "nil returned, but error expected"
	errNNil string = "error returned, but nil expected"
)

func TestErrorfNil(t *testing.T) {
	err := errorf(nil)
	if err == nil {
		t.Error(errNil)
	}
}

func TestNil(t *testing.T) {
	err := NilPtr()
	if err == nil {
		t.Fatal(errNil)
	}
	testValidJson(t, err)
	testEqualJson(t, err, &nilPtr)
}

func testValidJson(t *testing.T, e error) {
	if !json.Valid([]byte(fmt.Sprintf("%v", e))) {
		t.Error("not in valid json format")
	}
}

func testEqualJson(t *testing.T, e error, emsg *errmsg) {
	if emsg == nil {
		t.Fatal("nil pointer")
	}
	j, errj := json.Marshal(&errwrap{*emsg})
	if errj != nil {
		t.Fatal(errj)
	}
	if fmt.Sprintf("%v", e) != string(j) {
		t.Errorf("%v does not equal %v", e, string(j))
	}
}

func BenchmarkNil(b *testing.B) {
	var err [2]error
	// Reset benchmark timer
	b.ResetTimer()
	// Run benchmark with NilPtr()
	for i := 0; i < b.N; i++ {
		err[i&0x1] = NilPtr()
	}
}
