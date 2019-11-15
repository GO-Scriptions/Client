package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/190930-UTA-CW-Go/project-2/Client/web"
)

var remoteUser, remoteHost string

//Prescription is new and untested should be a struct that stores Prescription information
type Prescription struct {
	PRES []string
}

func main() {
	// Sets up a file server in current directory
	http.HandleFunc("/", Index)
	http.HandleFunc("/doctorlogin", DocLog)
	http.HandleFunc("/doctor", DocFunc)
	http.HandleFunc("/employeelogin", PhaLog)
	http.HandleFunc("/employee", PhaFunc)
	http.HandleFunc("/stock", Stock)
	http.HandleFunc("/prescription", Presc)
	http.HandleFunc("/docpres", Docpres)
	http.ListenAndServe(":80", nil)
}

//moved over from handlefunc.go ///////////////////////////////////

//LoginInfo structure to save your username and occupation
type LoginInfo struct {
	Username string
	Doctor   bool
}

var loginInfo = LoginInfo{}

// Index runs the index page
func Index(response http.ResponseWriter, request *http.Request) {
	loginInfo = LoginInfo{} //log out user if logged in
	temp, _ := template.ParseFiles("web/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
	connection := web.FirstConnect()
	fmt.Println(connection)
}

// DocLog HTTP Handler for Doctor Login
func DocLog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/doctorlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// DocFunc HTTP Handler for after Doctor logs in NEW AND UNTESTED
func DocFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/doctor.html")
	var cmd, dbResponse string
	//values of form text boxes
	uname := request.FormValue("uname")
	dpass := request.FormValue("dpass")

	// cmd0 cd into directory

	cmd = "/usr/local/go/bin/go run main.go --log d "
	cmd += uname
	cmd += " "
	cmd += dpass
	fmt.Println("command:", cmd)

	dbResponse = web.GenLogin(cmd)
	fmt.Println("db response:", dbResponse)

	if dbResponse == "true" {
		loginInfo.Doctor = true
		loginInfo.Username = uname // change http to doctor.html
	} else {
		loginInfo.Doctor = false
		loginInfo.Username = uname
	}
	temp.Execute(response, loginInfo)
}

//Docpres sends a new prescription to the database
func Docpres(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/docpres.html")
	var cmd, dbResponse string
	fname := request.FormValue("fname")
	lname := request.FormValue("lname")
	amount := request.FormValue("amount")
	prescription := request.FormValue("prescription")
	// go run main.go
	cmd += "/usr/local/go/bin/go run main.go --doc wp "
	// add command line arguments
	cmd += fname
	cmd += " "
	cmd += lname
	cmd += " "
	cmd += amount
	cmd += " "
	cmd += prescription
	fmt.Println("command:", cmd)
	// get database response
	dbResponse = web.ExecuteCommand(cmd)
	dbResponse = strings.TrimSpace(dbResponse)
	fmt.Println("db response:", dbResponse)
	temp.Execute(response, dbResponse) //changed loginInfo to dbResponse
}

// PhaLog HTTP Handler for Pharmasicst Login
func PhaLog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("employeelogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// PhaFunc HTTP Handler for after Pharmasicst logs in
func PhaFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("employee.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// Stock HTTP Hander for restocking the pharamacy
func Stock(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("stock.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// Presc HTTP Handler to view prescriptions ///////////massively changed and untested
func Presc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("prescription.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	var cmd, dbResponse string
	//values of form text boxes
	uname := request.FormValue("uname")
	// cmd0 cd into directory

	cmd = "/usr/local/go/bin/go run main.go --doc vp "
	cmd += uname
	fmt.Println("command:", cmd)

	dbResponse = web.ExecuteCommand(cmd)
	dbResponse = strings.TrimSpace(dbResponse)
	fmt.Println("dbResponse:", dbResponse)

	//changes datatype
	p := Prescription{PRES: make([]string, 1)}
	length := 0

	//adds dbResponse to struct Prescription line by line
	for l := 0; l < len(dbResponse); l = l + 1 {
		if dbResponse[l] != 10 {
			p.PRES[length] = p.PRES[length] + string(dbResponse[l])
		} else {
			p.PRES = append(p.PRES, "\n")
			length = length + 1
		}
	}

	temp.Execute(response, p)
}
