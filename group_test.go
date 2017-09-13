package grinder

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupCreation(t *testing.T) {
	g := New()

	group := g.Group("/group")

	assert.True(t, reflect.TypeOf(group).String() == "*grinder.Group")
}

func TestGroupGetRoute(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.GET("/path", func(c Context) error {
		return c.String(200, "This is a test")
	})

	found := g.router.getRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/group/path"]).String() == "grinder.Route")
}

func TestGroupPostRoute(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.POST("/path", func(c Context) error {
		return c.String(200, "This is a test")
	})

	found := g.router.getRoutes("POST")

	assert.True(t, reflect.TypeOf(found["POST/group/path"]).String() == "grinder.Route")
}

func TestGroupPatchRoute(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.PATCH("/path", func(c Context) error {
		return c.String(200, "This is a test")
	})

	found := g.router.getRoutes("PATCH")

	assert.True(t, reflect.TypeOf(found["PATCH/group/path"]).String() == "grinder.Route")
}

func TestGroupPutRoute(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.PUT("/path", func(c Context) error {
		return c.String(200, "This is a test")
	})

	found := g.router.getRoutes("PUT")

	assert.True(t, reflect.TypeOf(found["PUT/group/path"]).String() == "grinder.Route")
}

func TestGroupDeleteRoute(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.DELETE("/path", func(c Context) error {
		return c.String(200, "This is a test")
	})

	found := g.router.getRoutes("DELETE")

	assert.True(t, reflect.TypeOf(found["DELETE/group/path"]).String() == "grinder.Route")
}
