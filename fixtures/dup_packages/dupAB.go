package dup_packages // import "github.com/jibecompany/counterfeiter/v6/fixtures/dup_packages"

//counterfeiter:generate . DupAB
type DupAB interface {
	DupA
	DupB
}
