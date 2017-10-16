package middleware

import (
	"net/http"
	"net/http/httptest"
	// "reflect"
	"strings"
	"testing"

	"github.com/rinkbase/grinder"
	"github.com/stretchr/testify/assert"
)

func TestJWTError(t *testing.T) {
	g := grinder.New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()
	g.ServeHTTP(w, r)

	err := JWTError(g.GetContext())

	assert.Nil(t, err)
}
