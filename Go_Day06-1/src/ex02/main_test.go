package main_test

import (
	"context"
	"myArticles"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a main.App

func TestMain(t *testing.T) {
	a.Init(context.Background(), "postgres://golang_day06:1234@localhost:5432/golang_day06", main.Config{})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestMaxAmountOfRequests(t *testing.T) {
	numRequests := 200
	for i := 0; i < numRequests; i++ {
		req, _ := http.NewRequest("GET", "/", nil)
		response := executeRequest(req)
		if i >= 100 && response.Code == http.StatusTooManyRequests {
			return
		}

		if response.Code != http.StatusOK {
			t.Fatalf("Unexpected status code: %d, request num: %d", response.Code, i)
		}
	}

	t.Fatalf("rate limiting did not work")
}
