package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TesteHandlerGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 status code but got %d", resp.StatusCode)
	}
}
