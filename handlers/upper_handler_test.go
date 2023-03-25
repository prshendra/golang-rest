package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoutes(t *testing.T) {
	r := gin.Default()
	RegisterRoutes(r)

	// Create a JSON payload
	payload := []byte(`{"input": "abc"}`)

	req, err := http.NewRequest("POST", "/upper", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the request to the handler
	r.ServeHTTP(w, req)

	// Verify the response status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
	}

	// Verify the response body
	expected := `{"message":"ABC"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
