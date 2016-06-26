package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestHelloWorld_Ok(t *testing.T) {
	var response = httptest.NewRecorder()
	GetProductAll(response, nil)

	if response.Code != 200 {
		t.Errorf("Should have 200 status code. Response Code: %v", response.Code)
	}

	if response.Body.String() != "{ &#39;msg&#39;:&#39;Hello World&#39; }" {
		t.Errorf("Should have had a hello world. Body: %v", response.Body.String())
	}
}