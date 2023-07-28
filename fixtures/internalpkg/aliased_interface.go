package internalpkg

import "github.com/jibecompany/counterfeiter/v6/fixtures/internalpkg/internal"

//go:generate go run github.com/jibecompany/counterfeiter/v6 -generate
//counterfeiter:generate . Context

type Context = internal.Context
