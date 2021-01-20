package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"saas/router"
	"testing"
)

func TestIndexGetRouter(t *testing.T)  {
	engine := router.Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello world", w.Body.String())
}
