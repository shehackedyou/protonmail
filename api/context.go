package api

import (
	"context"
)

type pmapiContextKey string

const (
	retryContextKey = pmapiContextKey("retry")
	retryDisabled   = "disabled"

	authRefreshContextKey = pmapiContextKey("authRefresh")
	authRefreshDisabled   = "disabled"
)

func ContextWithoutRetry(parent context.Context) context.Context {
	return context.WithValue(parent, retryContextKey, retryDisabled)
}

func isRetryDisabled(ctx context.Context) bool {
	if v := ctx.Value(retryContextKey); v != nil {
		return v == retryDisabled
	}
	return false
}

func ContextWithoutAuthRefresh(parent context.Context) context.Context {
	return context.WithValue(parent, authRefreshContextKey, authRefreshDisabled)
}

func isAuthRefreshDisabled(ctx context.Context) bool {
	if v := ctx.Value(authRefreshContextKey); v != nil {
		return v == authRefreshDisabled
	}
	return false
}
