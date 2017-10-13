package middleware

import (
	"github.com/rinkbase/grinder"
	"net/http"
	"strings"
)

// CORSConfig configuration for CORS middleware
type CORSConfig struct {
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedHeaders []string `json:"allowed_headers"`
	ExposedHeaders []string `json:"exposed_headers"`
}

// DefaultCORSConfig handles the default CORS configuration for grinder
var DefaultCORSConfig = CORSConfig{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders: []string{"*"},
	ExposedHeaders: []string{"*"},
}

// CORSError returns a grinder Handler when an error is occured
func CORSError(ctx grinder.Context) error {
	return ctx.JSON(500, "CORS Error")
}

// CORS middleware for Grinder routes
func CORS(ctx grinder.Context, handler grinder.Handler) grinder.Handler {
	return CORSConfigured(ctx, handler, DefaultCORSConfig)
}

// CORSConfigured returns a configured CORS middleware
func CORSConfigured(ctx grinder.Context, handler grinder.Handler, config CORSConfig) grinder.Handler {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = DefaultCORSConfig.AllowedOrigins
	}

	if len(config.AllowedMethods) == 0 {
		config.AllowedMethods = DefaultCORSConfig.AllowedMethods
	}

	if len(config.AllowedHeaders) == 0 {
		config.AllowedHeaders = DefaultCORSConfig.AllowedHeaders
	}

	if len(config.ExposedHeaders) == 0 {
		config.ExposedHeaders = DefaultCORSConfig.ExposedHeaders
	}

	allowedOrigins := strings.Join(config.AllowedOrigins, ",")
	allowedMethods := strings.Join(config.AllowedMethods, ",")
	allowedHeaders := strings.Join(config.AllowedHeaders, ",")
	exposedHeaders := strings.Join(config.ExposedHeaders, ",")

	return func(ctx grinder.Context) error {
		req := ctx.Request()
		res := ctx.Response()

		if req.Method != "OPTIONS" {
			res.Header().Add("Vary", "Origin")
			res.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			if exposedHeaders != "" {
				res.Header().Set("Access-Control-Expose-Headers", exposedHeaders)
			}

			return handler(ctx)
		}

		res.Header().Add("Vary", "Origin")
		res.Header().Add("Vary", "Access-Control-Request-Method")
		res.Header().Add("Vary", "Access-Control-Request-Headers")
		res.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		res.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		res.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

		return ctx.Code(http.StatusNoContent)
	}
}
