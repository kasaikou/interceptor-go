package icept_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kasaikou/interceptor-go/icept"
)

func TestSmallSample_Interceptor(t *testing.T) {

	interceptor1 := icept.FromFnW[*[]int](func(ctx context.Context, writer *[]int, next icept.NextFnW[*[]int]) {
		*writer = append(*writer, 1)
		next(ctx, writer)
		*writer = append(*writer, 101)
	})

	interceptor2 := icept.FromFnW[*[]int](func(ctx context.Context, writer *[]int, next icept.NextFnW[*[]int]) {
		*writer = append(*writer, 2)
		next(ctx, writer)
		*writer = append(*writer, 102)
	})

	interceptor3 := icept.FromFnW[*[]int](func(ctx context.Context, writer *[]int, next icept.NextFnW[*[]int]) {
		*writer = append(*writer, 3)
		next(ctx, writer)
		*writer = append(*writer, 103)
	})

	interceptor := icept.New[struct{}, *[]int]().
		AppendW(interceptor1).
		AppendW(interceptor2).
		AppendW(interceptor3)

	entrypoint := interceptor.MakeTermination(func(ctx context.Context, reader struct{}, writer *[]int) {})

	result := make([]int, 0, 6)

	entrypoint(context.Background(), struct{}{}, &result)

	expected := []int{1, 2, 3, 103, 102, 101}
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("results is not unexpected value (want: %v, have: %v)", expected, result)
	}
}
