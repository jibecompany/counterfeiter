package fixtures

import (
	"github.com/jibecompany/counterfeiter/v6/fixtures/go-hyphenpackage"
)

//counterfeiter:generate . ImportsGoHyphenPackage
type ImportsGoHyphenPackage interface {
	UseHyphenType(hyphenpackage.HyphenType)
}
