package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandlerReturnsName(t *testing.T) {
	data, err := getBodyFromHandler(t, RootHandler)
	if err != nil {
		t.Errorf("unable to get body from handler: %s", err)
	}

	if !strings.Contains(string(data), "some-server") {
		t.Errorf("expected data to contain 'some-server' got %v", string(data))
	}
}

func getBodyFromHandler(t *testing.T, handler func(http.ResponseWriter, *http.Request)) ([]byte, error) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()

	handler(writer, req)
	res := writer.Result()
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
