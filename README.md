# tserr

[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/tserr)](https://goreportcard.com/report/github.com/thorstenrie/tserr)
[![CodeFactor](https://www.codefactor.io/repository/github/thorstenrie/tserr/badge)](https://www.codefactor.io/repository/github/thorstenrie/tserr)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorstenrie/tserr)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/tserr)](https://pkg.go.dev/mod/github.com/thorstenrie/tserr)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorstenrie/tserr)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorstenrie/tserr)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorstenrie/tserr)
![GitHub last commit](https://img.shields.io/github/last-commit/thorstenrie/tserr)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorstenrie/tserr)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorstenrie/tserr)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorstenrie/tserr)
![GitHub](https://img.shields.io/github/license/thorstenrie/tserr)

The package tserr is a simple error interface in [Go](https://go.dev/). The interface provides standardized error messages of type `error` by function calls. Errors may contain verbs, which are provided by the function arguments.

- **Simple**: Without configuration, just function calls
- **Easy to parse**: All error messages in JSON format.
- **Tested**: Unit tests with high [code coverage](https://gocover.io/github.com/thorstenrie/tserr)
- **Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std)

## Usage

In the Go app, the package is imported with

```
import "github.com/thorstenrie/tserr"
```

An error is called by a specific, corresponding function call, e.g., 

```
err1 := tserr.NilPtr()
```

The output of `fmt.Println(err1)` is

```
{"error":{"id":0,"code":500,"message":"nil pointer"}}
```

The error message may contain verbs to be filled by arguments, e.g., with one argument:

```
f := "foo.txt"
err2 := tserr.NotExistent(f)
```

The output of `fmt.Println(err2)` is

```
{"error":{"id":2,"code":404,"message":"foo.txt does not exist"}}
```

A function may hold multiple arguments used for more than one verb in the error message. Multiple arguments are passed to a function as a pointer to a struct, e.g.,

```
err3 := tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: "a", Y: "b"})
```

The output of `fmt.Println(err3)` is

```
{"error":{"id":6,"code":500,"message":"a does not equal b"}}
```

All error functions with multiple arguments first check, if the pointer to the argument struct is nil. If it is nil, the error function returns NilPtr. If the argument struct contains a field of type error, it is checked next. If the error field is nil, then the error function returns nil. Otherwise, it returns the corresponding error message.

## JSON format

The error messages are fomatted in the JSON format. The root element is named "error". Each error message has an "id" which is consecutively numbered. "code" is a relating HTTP status code. "message" contains the actual pre-defined error message.

```
{"error":{"id":<int>,"code":<int>,"message":"<string>"}}
```

## Example

```
package main

import (
	"fmt"

	"github.com/thorstenrie/tserr"
)

func main() {
	err1 := tserr.NilPtr()
	fmt.Println(err1)

	f := "foo.txt"
	err2 := tserr.NotExistent(f)
	fmt.Println(err2)

	err3 := tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: "a", Y: "b"})
	fmt.Println(err3)
}
```

[Run in Go Playground](https://go.dev/play/p/L5u1D2Iy_M5)

Output:
```
{"error":{"id":0,"code":500,"message":"nil pointer"}}
{"error":{"id":2,"code":404,"message":"foo.txt does not exist"}}
{"error":{"id":6,"code":400,"message":"a does not equal b"}}
```

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/tserr)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/tserr)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Ftserr)
