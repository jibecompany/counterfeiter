package dup_packages // import "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages"

import "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages/b/foo"

//counterfeiter:generate . DupB
type DupB interface {
	B() foo.S
}
