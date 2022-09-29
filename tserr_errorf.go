package tserr

import "fmt"

type errm struct {
	ec int
	m  string
}

func errorf(e errm, a ...any) error {
	return fmt.Errorf("%d: %w", e.ec, fmt.Errorf(e.m, a...))
}
