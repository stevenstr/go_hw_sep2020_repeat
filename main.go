/*
* Author: Stefan
* Last change: 20.09.2020
* Task: Basic http server
 */

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var message = []string{"Hello Golang World!"}

//GetFunc function
func GetFunc(w http.ResponseWriter, r *http.Request) {
	//set content type
	w.Header().Set("Content-Type", "application-json")
	//encode json
	json.NewEncoder(w).Encode(message)
}

func main() {
	//create router using mux
	router := mux.NewRouter()
	//create handler on /golang and call on it our GetFunc
	router.HandleFunc("/golang", GetFunc)
	//just listen and serve on 8000 port
	http.ListenAndServe(":8000", router)
}
