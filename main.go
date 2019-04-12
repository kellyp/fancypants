package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type hwResponse struct {
	Hello string `json:"hello"`
}

type dResponse struct {
	Answer int `json:"answer"`
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	resp := hwResponse{
		Hello: "world",
	}

	b, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("unable to marshal json"))
	}
	w.Write(b)
}

func double(w http.ResponseWriter, r *http.Request) {
	//ns := strings.TrimPrefix(r.RequestURI, "localhost:8080/double/")
	split := strings.Split(r.RequestURI, "/")
	ns := split[len(split)-1]
	n, err := strconv.Atoi(ns)
	if err != nil {
		w.Write([]byte("invalid integer"))
		return
	}

	resp := dResponse{
		Answer: n * 2,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("unable to marshal json"))
		return
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/helloworld", helloworld)
	http.HandleFunc("/double/", double)

	fmt.Println("server starting")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	fmt.Println("Ended with error: ", err)
}
