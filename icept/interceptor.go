package icept

import "context"

// InterceptorRWN is interceptor, it uses reader and writer.
// It is considerated patterns that input type (R, W) and output type (NR, NW) is different.
type InterceptorRWN[R, W, NR, NW any] struct{ interceptorRWN[R, W, NR, NW] }
type interceptorRWN[R, W, NR, NW any] interface {
	InterceptRW(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW])
}

func FromFnRWN[R, W, NR, NW any](fn func(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW])) InterceptorRWN[R, W, NR, NW] {

	if fn == nil {
		panic("fn is nil")
	}

	return InterceptorRWN[R, W, NR, NW]{interceptorRWN: interceptorFnRWN[R, W, NR, NW](fn)}
}

// AppendRW creates a new InterceptorRWN instance appended other interceptor, using reader and writer.
func (i InterceptorRWN[R, W, NR, NW]) AppendRW(interceptor InterceptorRW[NR, NW]) InterceptorRWN[R, W, NR, NW] {
	return FromFnRWN[R, W, NR, NW](func(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW]) {
		i.interceptorRWN.InterceptRW(ctx, reader, writer, func(ctx context.Context, reader NR, writer NW) {
			interceptor.interceptorRWN.InterceptRW(ctx, reader, writer, next)
		})
	})
}

func (i InterceptorRWN[R, W, NR, NW]) AppendW(interceptor InterceptorW[NW]) InterceptorRWN[R, W, NR, NW] {
	return FromFnRWN[R, W, NR, NW](func(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW]) {
		i.interceptorRWN.InterceptRW(ctx, reader, writer, func(ctx context.Context, reader NR, writer NW) {
			interceptor.interceptorWN.InterceptW(ctx, writer, func(ctx context.Context, writer NW) {
				next(ctx, reader, writer)
			})
		})
	})
}

// AppendR creates a new InterceptorRWN instance appended other interceptor using reader.
func (i InterceptorRWN[R, W, NR, NW]) AppendR(interceptor InterceptorR[NR]) InterceptorRWN[R, W, NR, NW] {
	return FromFnRWN[R, W, NR, NW](func(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW]) {
		i.interceptorRWN.InterceptRW(ctx, reader, writer, func(ctx context.Context, reader NR, writer NW) {
			interceptor.interceptorRN.InterceptR(ctx, reader, func(ctx context.Context, reader NR) {
				next(ctx, reader, writer)
			})
		})
	})
}

// Append creates a new InterceptorRWN instance appended other interceptor
func (i InterceptorRWN[R, W, NR, NW]) Append(interceptor Interceptor) InterceptorRWN[R, W, NR, NW] {
	return FromFnRWN[R, W, NR, NW](func(ctx context.Context, reader R, writer W, next NextFnRW[NR, NW]) {
		i.interceptorRWN.InterceptRW(ctx, reader, writer, func(ctx context.Context, reader NR, writer NW) {
			interceptor.interceptor.Intercept(ctx, func(ctx context.Context) {
				next(ctx, reader, writer)
			})
		})
	})
}

// MakeTermination makes a new Entrypoint with interceptor and termination function.
func (i InterceptorRWN[R, W, NR, NW]) MakeTermination(termination Termination[NR, NW]) Entrypoint[R, W] {
	return func(ctx context.Context, reader R, writer W) {
		i.interceptorRWN.InterceptRW(ctx, reader, writer, NextFnRW[NR, NW](termination))
	}
}

// InterceptorRWN is interceptor, it uses reader and writer.
type InterceptorRW[R, W any] InterceptorRWN[R, W, R, W]

func New[R, W any]() InterceptorRW[R, W] {
	return FromFnRW[R, W](func(ctx context.Context, reader R, writer W, next NextFnRW[R, W]) {
		next(ctx, reader, writer)
	})
}

func FromFnRW[R, W any](fn func(ctx context.Context, reader R, writer W, next NextFnRW[R, W])) InterceptorRW[R, W] {

	if fn == nil {
		panic("fn is nil")
	}

	return InterceptorRW[R, W]{interceptorRWN: interceptorFnRWN[R, W, R, W](fn)}
}

// RWN convertes from InterceptorRW to InterceptorRWN.
//
// it is short command for InterceptorRWN[R, W, R, W](interceptorRW)
func (i InterceptorRW[R, W]) RWN() InterceptorRWN[R, W, R, W] {
	return InterceptorRWN[R, W, R, W](i)
}

func (i InterceptorRW[R, W]) AppendRW(interceptor InterceptorRW[R, W]) InterceptorRW[R, W] {
	return InterceptorRW[R, W](i.RWN().AppendRW(interceptor))
}

func (i InterceptorRW[R, W]) AppendW(interceptor InterceptorW[W]) InterceptorRW[R, W] {
	return InterceptorRW[R, W]{i.RWN().AppendW(interceptor)}
}

func (i InterceptorRW[R, W]) AppendR(interceptor InterceptorR[R]) InterceptorRW[R, W] {
	return InterceptorRW[R, W](i.RWN().AppendR(interceptor))
}

func (i InterceptorRW[R, W]) Append(interceptor Interceptor) InterceptorRW[R, W] {
	return InterceptorRW[R, W](i.RWN().Append(interceptor))
}

func (i InterceptorRW[R, W]) MakeTermination(termination Termination[R, W]) Entrypoint[R, W] {
	return i.RWN().MakeTermination(termination)
}

type InterceptorWN[W, NW any] struct{ interceptorWN[W, NW] }
type interceptorWN[W, NW any] interface {
	InterceptW(ctx context.Context, writer W, next NextFnW[NW])
}

type InterceptorW[W any] InterceptorWN[W, W]

func FromFnW[W any](fn func(ctx context.Context, writer W, next NextFnW[W])) InterceptorW[W] {

	if fn == nil {
		panic("fn is nil")
	}

	return InterceptorW[W]{interceptorWN: interceptorFnWN[W, W](fn)}
}

type InterceptorRN[R, NR any] struct{ interceptorRN[R, NR] }
type interceptorRN[R, NR any] interface {
	InterceptR(ctx context.Context, reader R, next NextFnR[NR])
}

func FromFnRN[R, NR any](fn func(ctx context.Context, reader R, next NextFnR[NR])) InterceptorRN[R, NR] {

	if fn == nil {
		panic("fn is nil")
	}

	return InterceptorRN[R, NR]{interceptorRN: interceptorFnRN[R, NR](fn)}
}

type InterceptorR[R any] InterceptorRN[R, R]

func FromFnR[R any](fn func(ctx context.Context, reader R, next NextFnR[R])) InterceptorR[R] {

	if fn == nil {
		panic("fn is nil")
	}

	return InterceptorR[R]{interceptorRN: interceptorFnRN[R, R](fn)}
}

func (i InterceptorR[R]) RN() InterceptorRN[R, R] {
	return InterceptorRN[R, R](i)
}

type Interceptor struct{ interceptor }
type interceptor interface {
	Intercept(ctx context.Context, next NextFn)
}

func FromFn(fn func(ctx context.Context, next NextFn)) Interceptor {

	if fn == nil {
		panic("fn is nil")
	}

	return Interceptor{interceptor: interceptorFn(fn)}
}
