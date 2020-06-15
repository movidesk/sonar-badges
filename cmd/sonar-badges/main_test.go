package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestContains(t *testing.T) {
	arr := []string{
		"hello",
		"world",
	}

	got := contains(arr, "hello")

	if !got {
		t.Errorf("contains(arr, \"hello\") = %t; want true", got)
	}
}

func TestErrorResponse(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/test.svg", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"error\":\"not valid\"}", w.Body.String())
}
