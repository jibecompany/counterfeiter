package foo // import "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages/foo"

import (
	"github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages/a/foo"
	bfoo "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages/b/foo"
)

type S struct{}

//go:generate go run github.com/jibecompany/counterfeiter/v6 . MultiAB
type MultiAB interface {
	Mine() S
	foo.I
	bfoo.I
}
