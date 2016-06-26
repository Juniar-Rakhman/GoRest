package handlers

import (
	"net/http/httptest"
	"testing"
	"net/http"
)

//Coverage testing supported here. But there are still problem with mux.Vars.


func TestIndex_Ok(t *testing.T) {

	//ARRANGE
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	//ACT
	Index(w, r)

	//ASSERT
	if w.Code != 200 {
		t.Errorf("Should have 200 status code. Response Code: %v", w.Code)
	}

	if w.Body.String() != "{ &#39;msg&#39;:&#39;Welcome to GoRest&#39; }" {
		t.Errorf("Should have had a hello world. Body: %v", w.Body.String())
	}
}

func TestGetProductById_Ok(t *testing.T) {

	//ARRANGE
	r, _ := http.NewRequest("GET", "/product/1", nil)
	w := httptest.NewRecorder()

	GetProductById(w, r) //http://mrgossett.com/post/mux-vars-problem/

	//ASSERT

}