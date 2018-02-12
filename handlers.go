package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This worked")
}

func PiThingHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("things/thing.json")

	if err != nil {
		fmt.Println(err)
	}

	w.Write(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
