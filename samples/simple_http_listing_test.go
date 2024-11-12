package samples

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetGreet(t *testing.T) {
	req, err := http.NewRequest("GET", `http://127.0.0.1:8080/`, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Если хорошо работаем с response
	res := httptest.NewRecorder()
	GetGreet(res, req)

	exp := `<h1> Hello! im new web-server!</h1>`
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %v, go %v", exp, act)
	}

	resCode := res.Result().StatusCode

	if resCode != 200 {
		t.Fatalf("Expected status 200 got %v", resCode)
	}
}
