package grinder

import (
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	id   int
	name string
}

const JSON = `{"id":1,"name":"John Adams"}`

func TestContext(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	// Response
	assert.NotNil(t, c.Response())

	// Request
	assert.NotNil(t, c.Request())
}

func TestJSONResponse(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	err := c.JSON(200, user{1, "John Adams"})

	if assert.NoError(t, err) {
		assert.Equal(t, 200, w.Code)
	}
}

func TestJSONErrorResponse(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	err := c.JSON(200, math.Inf(1))

	assert.Error(t, err)
}

func TestStringResponse(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	err := c.String(200, "this is a test")

	if assert.NoError(t, err) {
		assert.Equal(t, 200, w.Code)
	}
}

func TestErrorResponse(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	err := c.HTTPError(500, "this is a test")

	if assert.NoError(t, err) {
		assert.Equal(t, 500, w.Code)
	}
}

func TestAddingParams(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	params := make(map[string]string)
	params["key1"] = "value"
	params["key2"] = "value"

	c.AddParams(params)

	assert.Equal(t, len(c.GetParams()), 2)
	assert.Equal(t, c.GetParam("key1"), "value")
}

func TestSettingHeaders(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)
	c.SetHeader("Content-Type", "text/html;charset=utf-8")

	assert.Equal(t, c.Response().Header().Get("Content-Type"), "text/html;charset=utf-8")
}

func TestGettingHeaders(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	c.Request().Header.Set("Content-Type", "text/html;charset=utf-8")

	assert.Equal(t, c.GetHeader("Content-Type"), "text/html;charset=utf-8")
}

func TestRedirect(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	err := c.Redirect(301, "/")

	if assert.NoError(t, err) {
		assert.Equal(t, 301, w.Code)
	}
}

func TestHasParam(t *testing.T) {
	g := New()

	r, _ := http.NewRequest("GET", "/uri?param=1", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := g.NewContext(w, r)

	assert.True(t, c.HasParam("param"))
}
