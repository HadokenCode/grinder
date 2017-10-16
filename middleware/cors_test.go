package middleware

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/rinkbase/grinder"
	"github.com/stretchr/testify/assert"
)

const JSON = `{"id":1,"name":"John Adams"}`

func TestCORS(t *testing.T) {
	g := grinder.New()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	ctx := g.NewContext(rec, req)
	nfh := grinder.NotFoundHandler(ctx)

	assert.True(t, reflect.TypeOf(nfh).String() == "*grinder.HTTPError")
}

func TestCORSError(t *testing.T) {
	g := grinder.New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()
	g.ServeHTTP(w, r)

	err := CORSError(g.GetContext())

	assert.Nil(t, err)
}
