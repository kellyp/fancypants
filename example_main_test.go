package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
)

func ExampleHelloworld() {
	req := httptest.NewRequest("GET", "localhost:8080/helloworld", nil)
	w := httptest.NewRecorder()
	helloworld(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	// Output:
	// 200
	// {"hello":"world"}
}
func ExampleHDouble() {
	req := httptest.NewRequest("GET", "localhost:8080/double/2", nil)
	w := httptest.NewRecorder()
	double(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	// Output:
	// 200
	// {"answer":4}
}
