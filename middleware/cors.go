package middleware

import (
	"github.com/rinkbase/grinder"
)

// CORSConfig configuration for CORS middleware
type CORSConfig struct {
	AllowOrigins   []string `json:"allow_origins"`
	AllowMethods   []string `json:"allow_methods"`
	AllowHeaders   []string `json:"allow_headers"`
	HeadersExposed []string `json:"headers_exposed"`
}

// DefaultCORSConfig handles the default CORS configuration for grinder
var DefaultCORSConfig = CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
}

// CORSError returns a grinder Handler when an error is occured
func CORSError(ctx grinder.Context) error {
	return ctx.JSON(500, "CORS Error")
}

// CORS middleware for Grinder routes
func CORS(ctx grinder.Context, handler grinder.Handler) grinder.Handler {
	return CORSError
}
