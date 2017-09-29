[![Build Status](https://travis-ci.org/rinkbase/grinder.svg?branch=master)](https://travis-ci.org/rinkbase/grinder)
[![Code Coverage](https://codecov.io/gh/rinkbase/grinder/branch/master/graph/badge.svg)](https://codecov.io/gh/rinkbase/grinder/branch/master/graph/badge.svg)

# Grinder
> Grinder (noun): a player that populates the lower lines or lower pairings. Has hands of stone, but is physical and works hard when he’s out on the ice. Usually beloved by the rest of the team.

Grinder is a Go based framework with the aim of making the development of microservices easier. It attempts to do all the hard work / heaving lifting, hence the name Grinder, so that you can focus on your service.

## Example
```
func main() {
	svc := grinder.New()

	// index handler
	svc.GET("/", func(ctx grinder.Context) error {
		return c.JSON(200, "Hello World!")
	})

	svc.Start()
}
```

## Routing

### Add Routes
```
// Add GET route
svc.GET('/endpoint', handler)

// Add POST route
svc.POST('/endpoint', handler)

// Add PATCH route
svc.PATCH('/endpoint', handler)

// Add PUT route
svc.PUT('/endpoint', handler)

// Add DELETE route
svc.DELETE('/endpoint', handler)
```

### Middleware

#### Included Middleware

JWT Middleware
```
// the JWT middleware gets the key/secret from the config which is set in .env
svc.GET("/endpoint", handler, JWT)
```

#### Creating Custom Middleware

To create custom middleware:
```
// Define middleware handler
middleware := func(ctx grinder.Context, handler grinder.Handler) grinder.Handler {
	// code here
	return handler
}

// Add middleware to route
svc.GET('/path', handler, middleware)
```

### Hooks

#### Before:
```
svc.Before(middleware1)
svc.Before(middleware2)
-- or --
svc.Before(middleware1, middleware2)
```

#### After:
```
svc.After(middleware1)
svc.After(middleware2)
-- or --
svc.After(middleware1, middleware2)
```

### Route Groups
```
// Create Route Group
g := svc.Group('/prefix')
g.GET('/endpoint', handler) // GET /prefix/endpoint
g.POST('/endpoint', handler) // POST /prefix/endpoint
```

Groups can have defined middleware. These middleware handlers will be executed for every route within the group:
```
g := svc.Group('/prefix', middlewareHandler1, middlewareHandler1)
```

If you need middleware to be executed on specific routes you can add middleware to the route definition:
```
g := svc.Group('/prefix')
g.GET('/endpoint', handler, middleware)
```
