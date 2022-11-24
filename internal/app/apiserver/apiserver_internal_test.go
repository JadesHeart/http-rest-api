package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test HandleHello with the help "testify", expected "Привет мир"
func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Привет, Мир!")
}
