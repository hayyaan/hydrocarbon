package xmux

import "golang.org/x/net/context"

// TestSetParamContext sets ParamHolder to context.Context for testing
func TestSetParamContext(ctx context.Context, p ParamHolder) context.Context {
	return context.WithValue(ctx, paramsKey, p)
}
