package grinder

import (
	"bytes"
	"regexp"
	"strings"
)

// Router struct holds all defined routes
type Router struct {
	routes map[string]map[string]Route
}

// Route struct holds all information about a defined route
type Route struct {
	method     string
	path       string
	handler    Handler
	middleware []Middleware
}

const pattern = `([aA-zZ0-9_-]+)`
const query = `[^&?]*?=[^&?]*`

// Add will add a new route to the Router.routes map
func (r *Router) Add(m string, p string, h Handler, f []Middleware) {
	// create new route
	route := Route{
		method:  m,
		path:    p,
		handler: h,
	}

	// add middleware handlers
	for _, v := range f {
		route.middleware = append(route.middleware, v)
	}

	// init routes map if its not been already
	if r.routes == nil {
		r.routes = make(map[string]map[string]Route)
	}

	if r.routes[m] == nil {
		r.routes[m] = make(map[string]Route)
	}

	r.routes[m][m+p] = route
}

// Handler as defined by Grinder
func (r *Route) Handler() Handler {
	return r.handler
}

// FindRoute searches the defined routes
func (r *Router) FindRoute(c Context) (bool, Route) {
	found := false
	var route Route // by default route is nil, i.e. Not Found

	method := c.Request().Method // requested method
	path := strings.Split(c.Request().URL.String(), "?")

	routes := r.getRoutes(method)

	if len(routes) != 0 {
		for k, v := range routes {
			formatted := format(k)

			re := regexp.MustCompile(`^` + formatted + `/?$`)

			if re.MatchString(method + path[0]) {
				found = true
				route = v

				// get URL params
				c.AddParams(r.parseURLParams(method, path[0], formatted, k))

				// get form params
				c.Request().ParseForm()
				c.AddParams(parseFormParams(c.Request().Form))

				if len(path) > 1 {
					c.AddParams(parseQueryParams(path[1]))
				}
			}
		}
	}

	return found, route
}

func (r *Router) parseURLParams(method string, url string, path string, route string) map[string]string {
	if path == "/" {
		return make(map[string]string)
	}

	// map of params to be returned
	params := make(map[string]string)

	// key regular expression (kre)
	kre := regexp.MustCompile(`:` + pattern)
	keys := kre.FindAllStringSubmatch(route, -1)

	// value regular express (vre)
	vre := regexp.MustCompile(path)
	values := vre.FindAllStringSubmatch(method+url, -1)[0][1:]

	// assign keys to values
	for i, v := range keys {
		params[v[1]] = values[i]
	}

	return params
}

func parseQueryParams(url string) map[string]string {
	params := make(map[string]string)

	qre := regexp.MustCompile(query)
	q := qre.FindAllStringSubmatch(url, -1)

	for _, query := range q {
		values := strings.Split(query[0], "=")
		params[values[0]] = values[1]
	}

	params["param"] = "1"
	return params
}

func parseFormParams(form map[string][]string) map[string]string {
	params := make(map[string]string)

	for k, v := range form {
		params[k] = v[0]
	}

	return params
}

func (r *Router) getRoutes(method string) map[string]Route {
	route := make(map[string]Route)

	if val, exists := r.routes[method]; exists {
		return val
	}

	return route
}

func format(route string) string {
	var formatted bytes.Buffer

	re := regexp.MustCompile(`:` + pattern)

	formatted.WriteString(re.ReplaceAllString(route, pattern))

	return formatted.String()
}
