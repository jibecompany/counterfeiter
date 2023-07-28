package sync // import "github.com/jibecompany/counterfeiter/v6/fixtures/sync"

//go:generate go run github.com/jibecompany/counterfeiter/v6 . SyncSomething
type SyncSomething interface {
	DoThings(string, uint64) (int, error)
	DoNothing()
	DoASlice([]byte)
	DoAnArray([4]byte)
}
