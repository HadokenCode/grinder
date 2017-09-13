package grinder

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingRoute(t *testing.T) {
	g := New()

	g.GET("/path", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/path"]).String() == "grinder.Route")
}

func TestAddingRouteWithMiddleware(t *testing.T) {
	g := New()

	g.GET("/path", func(c Context) error {
		return c.JSON(200, "This is a test")
	}, middlware)

	found := g.router.getRoutes("GET")

	assert.Equal(t, 1, len(found["GET/path"].middleware))
}

func TestHandlerReturnsHandler(t *testing.T) {
	g := New()

	g.GET("/path", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("GET")
	route := found["GET/path"]

	assert.True(t, reflect.TypeOf(route.Handler()).String() == "grinder.Handler")
}

func TestGetRoutesReturnsRoutes(t *testing.T) {
	g := New()

	g.GET("/path", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/path"]).String() == "grinder.Route")
}

func TestGetRoutesReturnsEmptyIfRouteNotFound(t *testing.T) {
	g := New()

	g.GET("/path", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("POST")

	assert.Empty(t, found)
}
