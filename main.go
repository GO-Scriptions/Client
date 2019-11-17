package main

import (
	"net/http"
	"html/template"
	"fmt"
	"strings"
	
	"github.com/project-2/Client/web"
)

//LoginInfo structure to save your username and occupation
type LoginInfo struct {
	Username string
	Doctor   bool
}
type Prescription struct {
	PRES []string
}
var loginInfo = LoginInfo{}

// Index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, loginInfo)
}

// DocLog HTTP Handler for Doctor Login
func DocLog(response http.ResponseWriter, request *http.Request) { 

	temp, _ := template.ParseFiles("web/doctorlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, loginInfo)
}

func logout(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/index.html")

    	loginInfo = LoginInfo{}
	temp.Execute(response, nil)
}
// DocFunc HTTP Handler for after Doctor logs in NEW AND UNTESTED
func DocFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/doctor.html")
	var cmd, dbResponse string

	// if not logged in
	if !loginInfo.Doctor {
		//values of form text boxes
        	uname := request.FormValue("uname")
        	dpass := request.FormValue("dpass")

        	// cmd0 cd into directory
		cmd =  "cd go/src/github.com/Database;"
        	cmd += "/usr/local/go/bin/go run main.go --log d "
        	cmd += uname
        	cmd += " "
        	cmd += dpass
        	fmt.Println("command:", cmd)
		dbResponse = web.ExecuteCommand(cmd)
        	fmt.Println("db response:", dbResponse)
		// convert response to array of words
		words := strings.Fields(dbResponse)

        	if words[0] == "true" {
                	loginInfo.Doctor = true
                	loginInfo.Username = uname // change http to doctor.html
        	} else {
			fmt.Println("invalid login")
        	}
	}
	temp.Execute(response, loginInfo)
}

//Docpres sends a new prescription to the database
func Docpres(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/docpres.html")
	var cmd, db_response string
	fname := request.FormValue("fname")
	lname := request.FormValue("lname")
	amount := request.FormValue("amount")
	prescription := request.FormValue("prescription")
	// go run main.go
	cmd =  "cd go/src/github.com/Database;"
	cmd += "/usr/local/go/bin/go run main.go --doc wp "
	// add command line arguments
	cmd += loginInfo.Username //loginInfo uname
	cmd += " "
	cmd += fname
	cmd += " "
	cmd += lname
	cmd += " "
	cmd += amount
	cmd += " "
	cmd += prescription
	fmt.Println("command:", cmd)
	// get database response
	db_response = web.ExecuteCommand(cmd)
	db_response = strings.TrimSpace(db_response)
	fmt.Println("db response:", db_response)
	temp.Execute(response, loginInfo)
}

// Presc HTTP Handler to view prescriptions ///////////massively changed and untested
func Presc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("web/prescription.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	var cmd, dbResponse string
	//var dbResponse string
	//values of form text boxes
	uname := loginInfo.Username
	// cmd0 cd into directory
	
	cmd =  "cd go/src/github.com/Database;"
	cmd += "/usr/local/go/bin/go run main.go --doc vp "
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

	//temp.Execute(response, p)
	temp.Execute(response, p)
}

func main() {
	connection := web.FirstConnect()
        fmt.Println(connection)

	// Sets up a file server in current directory
	http.HandleFunc("/", index)
	http.HandleFunc("/doctorlogin", DocLog)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/doctor", DocFunc)
	http.HandleFunc("/docpres", Docpres)
	http.HandleFunc("/prescription", Presc)
	http.ListenAndServe(":80", nil)
}
