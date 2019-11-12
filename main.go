package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var instanceURL = "http://ec2-18-224-140-238.us-east-2.compute.amazonaws.com/"
var fromRemote = "Unchanged from remote"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", getEvent)
	router.HandleFunc("/response", gotEvent)
	log.Fatal(http.ListenAndServe(":80", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home!")
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting an Event!")
	req, er0 := http.NewRequest("GET", instanceURL, nil)
	if er0 != nil {
		log.Fatal(er0)
	}
	res, er1 := http.DefaultClient.Do(req)
	if er1 != nil {
		log.Fatal(er1)
	}

	defer res.Body.Close()
	fromRemote, _ := ioutil.ReadAll(res.Body)

	log.Println(fromRemote)
}

func gotEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fromRemote)
}
