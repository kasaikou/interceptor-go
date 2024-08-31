package icept

import "context"

type NextFn func(ctx context.Context)

type NextFnR[R any] func(ctx context.Context, reader R)

type NextFnW[W any] func(ctx context.Context, writer W)

type NextFnRW[R, W any] func(ctx context.Context, reader R, writer W)
