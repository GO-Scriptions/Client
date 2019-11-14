package web

import (
	"fmt"
	"net/http"
	"os/exec"
	"text/template"
)

// Index runs the index page
func Index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
	connection := firstConnect()
	fmt.Println(connection)
}

// DocLog HTTP Handler for Doctor Login
func DocLog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/doctorlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
	//values of form text boxes
	fname := request.FormValue("fname")
	lname := request.FormValue("lname")
	dpass := request.FormValue("dpass")
	//dsubmit := request.FormValue("dsubmit")		//don't know how to get it to run only when submit is pressed
	if fname == "Bob" && lname == "Builder" && dpass == "yes" { //change to compare to database
		fmt.Println("Welcome", fname, lname)   //just for testing
		fmt.Println("Your password is", dpass) //just for testing
		//fmt.Println(dsubmit)

		//add flags to make it check the Doctors table for user info in this case --doc........will need to change this code for other machines
		syst, err := exec.Command("ssh", "user1@192.168.56.102", "cd", "test", ";", "go", "run", "main.go", "--doc").Output()
		if err != nil {
			fmt.Println(err)
		}
		result := string(syst)
		fmt.Println(result)
	} else {
		fmt.Println("wrong") //just for testing
		//fmt.Println(dsubmit)

		//trying to print out when invalid info is entered
		//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		//err = t.ExecuteTemplate(out, "T", "<script>alert('your information was incorrect please try again')</script>")
	}
}

// DocFunc HTTP Handler for after Doctor logs in
func DocFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/doctor.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// PatLog HTTP Handler for Patient Login
func PatLog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/patientlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// PatFunc HTTP Handler for after Patient logs in
func PatFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/patient.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// PhaLog HTTP Handler for Pharmasicst Login
func PhaLog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/employeelogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// PhaFunc HTTP Handler for after Pharmasicst logs in
func PhaFunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/employee.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// Stock HTTP Hander for restocking the pharamacy
func Stock(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/stock.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

// Presc HTTP Handler for Pharmasicst to view prescriptions
func Presc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/prescription.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}
