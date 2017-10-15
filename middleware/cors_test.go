package middleware

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/rinkbase/grinder"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	g := grinder.New()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	ctx := g.NewContext(rec, req)
	nfh := grinder.NotFoundHandler(ctx)

	assert.True(t, reflect.TypeOf(nfh).String() == "*grinder.HTTPError")
}
