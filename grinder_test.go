package grinder

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedGrinder struct {
	mock.Mock
}

func handler(c Context) error {
	return nil
}

func middlware(c Context, handler Handler) Handler {
	return handler
}

func TestNewGrinder(t *testing.T) {
	g := New()
	assert.NotNil(t, g)
}

func TestRouterIsGrinderRouter(t *testing.T) {
	g := New()
	assert.True(t, reflect.TypeOf(g.router).String() == "*grinder.Router")
}

func TestAddBeforeMiddleware(t *testing.T) {
	g := New()

	mw1 := func(c Context, handler Handler) Handler {
		return handler
	}

	mw2 := func(c Context, handler Handler) Handler {
		return handler
	}

	g.Before(mw1, mw2)

	assert.True(t, len(g.before) == 2)
}

func TestAddAfterMiddleware(t *testing.T) {
	g := New()

	mw1 := func(c Context, handler Handler) Handler {
		return handler
	}

	mw2 := func(c Context, handler Handler) Handler {
		return handler
	}

	g.After(mw1, mw2)

	assert.True(t, len(g.after) == 2)
}

func TestCreateGroup(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.GET("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("GET")

	assert.True(t, len(found) == 1)
}

func TestGroupRouteAddedCorrectly(t *testing.T) {
	g := New()

	group := g.Group("/group")
	group.GET("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/group/test"]).String() == "grinder.Route")
}

func TestGroupRouteMiddlewareAddedCorrectly(t *testing.T) {
	g := New()

	group := g.Group("/group", middlware)
	group.GET("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	assert.Equal(t, len(group.middleware), 1)
}

func TestAddGetRoute(t *testing.T) {
	g := New()

	g.GET("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/test"]).String() == "grinder.Route")
}

func TestAddPostRoute(t *testing.T) {
	g := New()

	g.POST("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("POST")

	assert.True(t, reflect.TypeOf(found["POST/test"]).String() == "grinder.Route")
}

func TestConfigLoad(t *testing.T) {
	// config := config.Load("./testdata")
	var config map[string]string
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	assert.True(t, reflect.TypeOf(config).String() == "map[string]string")
}

func TestAddPatchRoute(t *testing.T) {
	g := New()

	g.PATCH("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("PATCH")

	assert.True(t, reflect.TypeOf(found["PATCH/test"]).String() == "grinder.Route")
}

func TestAddPutRoute(t *testing.T) {
	g := New()

	g.PUT("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("PUT")

	assert.True(t, reflect.TypeOf(found["PUT/test"]).String() == "grinder.Route")
}

func TestAddDeleteRoute(t *testing.T) {
	g := New()

	g.DELETE("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("DELETE")

	assert.True(t, reflect.TypeOf(found["DELETE/test"]).String() == "grinder.Route")
}

func TestAddOptionsRoute(t *testing.T) {
	g := New()

	g.OPTIONS("/test", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	found := g.router.getRoutes("OPTIONS")

	assert.True(t, reflect.TypeOf(found["OPTIONS/test"]).String() == "grinder.Route")
}

func TestGetContext(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader("JSON"))
	w := httptest.NewRecorder()

	g.context = g.NewContext(w, r)

	assert.True(t, reflect.TypeOf(g.GetContext()).String() == "*grinder.context")
}

func TestNotFoundHandler(t *testing.T) {
	g := New()

	g.GET("/", func(c Context) error {
		return c.JSON(200, "This is a test")
	})

	r, _ := http.NewRequest("GET", "/blah", strings.NewReader("String"))
	w := httptest.NewRecorder()

	g.ServeHTTP(w, r)

	if assert.Equal(t, 404, w.Code) {
		assert.Equal(t, "\"Not Found\"", w.Body.String())
	}
}

func TestErrorHTTP(t *testing.T) {
	err := &HTTPError{
		Code:    404,
		Message: "Not Found",
	}

	assert.Equal(t, "code=404, message=Not Found", err.Error())
}
