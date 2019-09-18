package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {

	var request *http.Request
	var err error
	request, err = http.NewRequest("GET", "/test/this/path", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EchoHandler)
	handler.ServeHTTP(recorder, request)

	// validate status
	if recorder.Code != http.StatusOK {
		t.Error("Error: expected ok status")
	}

	// unmarshal body
	var echoData EchoData
	err = json.Unmarshal(recorder.Body.Bytes(), &echoData)
	if err != nil {
		t.Error("Error: failed to unmarshal JSON")
	}

	// validate method
	if echoData.Method != http.MethodGet {
		t.Errorf("Error: expected GET: received: %v", echoData.Method)
	}

	// validate path
	if echoData.Path != "/test/this/path" {
		t.Errorf("Error: invalid path: expected /test/this/path: received: %v", echoData.Path)
	}

}
