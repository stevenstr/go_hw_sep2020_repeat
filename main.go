/*
* Author: Stefan
* Last change: 20.09.2020
* Task: Basic http server
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message = []string{"Hello Golang World!"}

//GetURIFunc function
func GetURIFunc(w http.ResponseWriter, r *http.Request) {
	//set content type
	w.Header().Set("Content-Type", "application-json")
	//encode json
	err := json.NewEncoder(w).Encode(message)
	//cover the errors
	if err != nil {
		fmt.Println("Error:", err)
	}
}

//GetURIFmt function
func GetURIFmt(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Golang World!!!")
	//cover the errors
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	//create router using mux
	router := mux.NewRouter()
	//create handler on /golang and call on it our GetFunc
	router.HandleFunc("/golang", GetURIFunc)
	router.HandleFunc("/golangx", GetURIFmt)
	//just listen and serve on 8000 port
	http.ListenAndServe(":8000", router)
}
