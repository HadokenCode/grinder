package grinder

// Group holds information about group route
type Group struct {
	prefix     string
	middleware []Middleware
	grinder    *Grinder
}

// GET adds a HTTP GET method to the group
func (g *Group) GET(e string, f Handler, m ...Middleware) {
	g.add("GET", e, f, m...)
}

// POST adds a HTTP POST method to the group
func (g *Group) POST(e string, f Handler, m ...Middleware) {
	g.add("POST", e, f, m...)
}

// PATCH adds a HTTP PATCH method to the group
func (g *Group) PATCH(e string, f Handler, m ...Middleware) {
	g.add("PATCH", e, f, m...)
}

// PUT adds a HTTP PUT method to the group
func (g *Group) PUT(e string, f Handler, m ...Middleware) {
	g.add("PUT", e, f, m...)
}

// DELETE adds a HTTP DELETE method to the group
func (g *Group) DELETE(e string, f Handler, m ...Middleware) {
	g.add("DELETE", e, f, m...)
}

func (g *Group) add(method string, e string, h Handler, middleware ...Middleware) {
	m := []Middleware{}
	m = append(m, g.middleware...)
	m = append(m, middleware...)
	g.grinder.add(method, g.prefix+e, h, m)
}
