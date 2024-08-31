package icept

import "context"

type interceptorFnRWN[R, W, NR, NW any] func(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW])

func (fn interceptorFnRWN[R, W, NR, NW]) InterceptRW(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW]) {
	fn(ctx, reader, writer, next)
}

type interceptorFnWN[W, NW any] func(ctx context.Context, writer W, next NextFnW[NW])

func (fn interceptorFnWN[W, NW]) InterceptW(ctx context.Context, writer W, next NextFnW[NW]) {
	fn(ctx, writer, next)
}

type interceptorFnRN[R, NR any] func(ctx context.Context, reader R, next NextFnR[NR])

func (fn interceptorFnRN[R, NR]) InterceptR(ctx context.Context, reader R, next NextFnR[NR]) {
	fn(ctx, reader, next)
}

type interceptorFn func(ctx context.Context, next NextFn)

func (fn interceptorFn) Intercept(ctx context.Context, next NextFn) {
	fn(ctx, next)
}
