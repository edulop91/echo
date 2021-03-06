package web

import (
	"context"
	"time"
)

// Cancel injects the context for a given action with a cancel func.
func Cancel(action Action) Action {
	return func(ctx *Ctx) Result {
		ctx.ctx, ctx.cancel = context.WithCancel(context.Background())
		return action(ctx)
	}
}

// Timeout injects the context for a given action with a timeout context.
func Timeout(d time.Duration) Middleware {
	return func(action Action) Action {
		return func(ctx *Ctx) Result {
			ctx.ctx, ctx.cancel = context.WithTimeout(context.Background(), d)
			return action(ctx)
		}
	}
}

// ViewProviderAsDefault sets the context.CurrrentProvider() equal to context.View().
func ViewProviderAsDefault(action Action) Action {
	return func(ctx *Ctx) Result {
		return action(ctx.WithDefaultResultProvider(ctx.View()))
	}
}

// JSONProviderAsDefault sets the context.CurrrentProvider() equal to context.API().
func JSONProviderAsDefault(action Action) Action {
	return func(ctx *Ctx) Result {
		return action(ctx.WithDefaultResultProvider(ctx.JSON()))
	}
}

// XMLProviderAsDefault sets the context.CurrrentProvider() equal to context.API().
func XMLProviderAsDefault(action Action) Action {
	return func(ctx *Ctx) Result {
		return action(ctx.WithDefaultResultProvider(ctx.XML()))
	}
}
