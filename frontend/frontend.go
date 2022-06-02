package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, req *http.Request) {
	rsp, err := http.Get("http://127.0.0.1:8090/")
	if err != nil {
		fmt.Printf("Backend server query error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("Can't read backend server response body error: %s", err)
		writeInternalErrorResponse(w)
		return
	}
	rspText := fmt.Sprintf("Message from backend '%s'", strings.TrimSuffix(string(body), "\n"))
	_, err = fmt.Fprintf(w, rspText)
	if err != nil {
		fmt.Printf("Response produce error: %s", err)
		writeInternalErrorResponse(w)
	}
}

func writeInternalErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		fmt.Println(err)
	}
}
