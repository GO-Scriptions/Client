package web

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

//LoginInfo structure to save your username and occupation
type LoginInfo struct {
	Username string
	Doctor   bool
}

var loginInfo = LoginInfo{}

// Index runs the index page
func Index(response http.ResponseWriter, request *http.Request) {
	loginInfo = LoginInfo{} //log out user if logged in
	temp, _ := template.ParseFiles("index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
	connection := firstConnect()
	fmt.Println(connection)
}

// DocLog HTTP Handler for Doctor Login
func DocLog(response http.ResponseWriter, request *http.Request) {
	var cmd, dbResponse string

	temp, _ := template.ParseFiles("doctorlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)

	//values of form text boxes
	uname := request.FormValue("uname")
	dpass := request.FormValue("dpass")

	// cmd0 cd into directory

	cmd = "/usr/local/go/bin/go run main.go --log d"
	cmd += uname
	cmd += " "
	cmd += dpass
	fmt.Println("command:", cmd)

	dbResponse = genLogin(cmd)
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

// DocFunc HTTP Handler for after Doctor logs in NEW AND UNTESTED
func DocFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("doctor.html")
	var cmd, db_response string
	fname := request.FormValue("fname")
	lname := request.FormValue("lname")
	amount := request.FormValue("amount")
	prescription := request.FormValue("prescription")
	// go run main.go
	cmd += "/usr/local/go/bin/go run main.go --doc wp"
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
	db_response = ExecuteCommand(cmd)
	db_response = strings.TrimSpace(db_response)
	fmt.Println("db response:", db_response)
	temp.Execute(response, loginInfo)
}

// PatLog HTTP Handler for Patient Login
func PatLog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("patientlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// PatFunc HTTP Handler for after Patient logs in
func PatFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("patient.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
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

// Presc HTTP Handler for Pharmasicst to view prescriptions
func Presc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("prescription.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}
