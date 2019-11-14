package main

import (
	"net/http"

	"github.com/GO-Scriptions/Client/web"
)

var remoteUser, remoteHost string

func main() {
	// Sets up a file server in current directory
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/doctorlogin", web.DocLog)
	http.HandleFunc("/doctor", web.DocFunc)
	http.HandleFunc("/patientlogin", web.PatLog)
	http.HandleFunc("/patient", web.PatFunc)
	http.HandleFunc("/employeelogin", web.PhaLog)
	http.HandleFunc("/employee", web.PhaFunc)
	http.HandleFunc("/stock", web.Stock)
	http.HandleFunc("/prescription", web.Presc)
	http.ListenAndServe(":80", nil)
}
