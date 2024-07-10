package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPage(t *testing.T) {
	t.Run("returns hello world", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		Index(res, req)

		pageResponse := res.Body.String()
		expected := "Hello, World!"

		if pageResponse != expected {
			t.Errorf("Expected %s but got %s", expected, pageResponse)
		}
	})
}
