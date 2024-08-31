package icept

import "context"

type EmptyInterceptor struct{}

func (EmptyInterceptor) Call(ctx context.Context, reader struct{}, writer struct{}, next NextFnRW[struct{}, struct{}]) {
	next(ctx, reader, writer)
}
