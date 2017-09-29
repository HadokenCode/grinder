package middleware

import (
	"github.com/rinkbase/grinder"
)

// CORSError returns a grinder Handler when an error is occured
func CORSError(ctx grinder.Context) error {
	return ctx.JSON(500, "CORS Error")
}

// CORS middleware for Grinder routes
func CORS(ctx grinder.Context, handler grinder.Handler) grinder.Handler {
	return CORSError
}
