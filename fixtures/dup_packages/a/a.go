package a // import "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages/a"

import "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages/a/foo"

//go:generate go run github.com/jibecompany/counterfeiter/v6 . A
type A interface {
	V1() foo.I
}
