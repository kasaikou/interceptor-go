package interceptorgo

import (
	"context"
	"fmt"

	"github.com/kasaikou/interceptor-go/icept"
)

func Example() {

	interceptor1 := icept.FromFnRW[struct{}, struct{}](func(ctx context.Context, reader, writer struct{}, next icept.NextFnRW[struct{}, struct{}]) {
		fmt.Println("Start interceptor1")
		next(ctx, reader, writer)
		fmt.Println("End interceptor1")
	})
	interceptor2 := icept.FromFn(func(ctx context.Context, next icept.NextFn) {
		fmt.Println("Start interceptor2")
		next(ctx)
		fmt.Println("End interceptor2")
	})
	interceptor3 := icept.FromFn(func(ctx context.Context, next icept.NextFn) {
		fmt.Println("Start interceptor3")
		next(ctx)
		fmt.Println("End interceptor3")
	})

	interceptor := interceptor1.Append(interceptor2).Append(interceptor3)
	entrypoint := interceptor.MakeTermination(func(ctx context.Context, reader, writer struct{}) {
		fmt.Println("Call workload")
	})

	entrypoint(context.Background(), struct{}{}, struct{}{})
}
