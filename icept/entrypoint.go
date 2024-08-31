package icept

import "context"

type Termination[R, W any] func(ctx context.Context, reader R, writer W)

type Entrypoint[R, W any] func(ctx context.Context, reader R, writer W)

// func ConvertRW[IR, IW, OR, OW, CR, CW any](interceptor InterceptorRWN[IR, IW, OR, OW], converter InterceptorRWN[OR, OW, CR, CW]) InterceptorRWN[IR, IW, CR, CW] {

// }

// func ConvertR[IR, IW, OR, OW, CR any](interceptor InterceptorRWN[IR, IW, OR, OW], converter InterceptorRN[OR, CR]) InterceptorRWN[IR, IW, CR, OW] {

// }
