package grinder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
)

// Grinder struct holds router and context for framework
type Grinder struct {
	context Context
	router  *Router
	after   []Middleware
	before  []Middleware
}

// Handler basic function to router handlers
type Handler func(Context) error

// Middleware defines a function to process middleware
type Middleware func(Context, Handler) Handler

// NotFoundHandler default 404 handler for not found routes
var NotFoundHandler = func(c Context) error {
	b, _ := json.Marshal("Not Found")

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(404)
	c.Response().Write([]byte(b))

	return NewHTTPError(http.StatusNotFound)
}

// HTTPError handles structure of new HTTP error
type HTTPError struct {
	Code    int
	Message interface{}
	Inner   error
}

// NewHTTPError creates new HTTP error
func NewHTTPError(code int, message ...interface{}) *HTTPError {
	err := &HTTPError{
		Code:    code,
		Message: http.StatusText(code),
	}

	return err
}

// makes HTTPError adhere to the Error interface
func (e *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", e.Code, e.Message)
}

// New creates new Grinder instance
func New() *Grinder {
	// return Grinder struct
	return &Grinder{
		router: new(Router),
	}
}

// Before adds a middleware function to be executed before the route handler
func (g *Grinder) Before(m ...Middleware) {
	for i := 0; i < len(m); i++ {
		g.before = append(g.before, m[i])
	}
}

// After adds a middleware function to be executed after the route handler
func (g *Grinder) After(m ...Middleware) {
	for i := 0; i < len(m); i++ {
		g.after = append(g.after, m[i])
	}
}

// GetContext returns current context
func (g *Grinder) GetContext() Context {
	return g.context
}

// Group creates a route group with common prefix
func (g *Grinder) Group(prefix string, middleware ...Middleware) *Group {
	group := &Group{prefix: prefix, grinder: g}

	// add middleware handlers
	for _, v := range middleware {
		group.middleware = append(group.middleware, v)
	}

	return group
}

// GET adds a HTTP GET route to router
func (g *Grinder) GET(e string, f Handler, m ...Middleware) {
	g.add("GET", e, f, m)
}

// POST adds a HTTP POST route to router
func (g *Grinder) POST(e string, f Handler, m ...Middleware) {
	g.add("POST", e, f, m)
}

// PATCH adds a HTTP PATCH route to router
func (g *Grinder) PATCH(e string, f Handler, m ...Middleware) {
	g.add("PATCH", e, f, m)
}

// PUT adds a HTTP PUT route to router
func (g *Grinder) PUT(e string, f Handler, m ...Middleware) {
	g.add("PUT", e, f, m)
}

// DELETE adds a HTTP DELETE route to router
func (g *Grinder) DELETE(e string, f Handler, m ...Middleware) {
	g.add("DELETE", e, f, m)
}

// OPTIONS adds a HTTP OPTIONS route to router
func (g *Grinder) OPTIONS(e string, f Handler, m ...Middleware) {
	g.add("OPTIONS", e, f, m)
}

// Start initates the framework to start listening for requests
func (g *Grinder) Start() {
	server := g.NewServer()
	server.ListenAndServe()
}

// NewServer creates a server object for services
func (g *Grinder) NewServer() http.Server {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := http.Server{
		Addr:    config["PORT"],
		Handler: g,
	}

	return server
}

func (g *Grinder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.context = g.NewContext(w, r)

	if found, route := g.router.FindRoute(g.context); found != false {
		// execute middleware chains
		handler := func(c Context) error {
			handler := route.Handler()

			// execute middleware chain
			for i := 0; i < len(route.middleware); i++ {
				handler = route.middleware[i](c, handler)
			}

			return handler(g.context)
		}

		// execute before middleware
		if len(g.before) > 0 {
			for i := 0; i < len(g.before); i++ {
				handler = g.before[i](g.context, handler)
			}
		}

		// Execute chain
		if err := handler(g.context); err != nil {
			panic(err)
		}

		// execute after middleware
		if len(g.after) > 0 {
			for i := 0; i < len(g.after); i++ {
				handler = g.after[i](g.context, handler)
			}
		}

		return
	}

	// Route was not found
	NotFoundHandler(g.context)
	return
}

// NewContext creates a fresh context for framework
func (g *Grinder) NewContext(w http.ResponseWriter, r *http.Request) Context {
	return &context{
		request:  r,
		response: NewResponse(w),
	}
}

func (g *Grinder) add(a string, e string, f Handler, m []Middleware) {
	p := strings.Split(e, "?")
	g.router.Add(a, p[0], func(c Context) error {
		fn := f
		return fn(c)
	}, m)
}
