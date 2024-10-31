package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/store/sqlstore/teststore"

	"github.com/stretchr/testify/assert"
)

func TestServert_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)

}
