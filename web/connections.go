package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

var remoteUser, remoteHost string

var port = "22"

func FirstConnect() string {
	var status string
        fmt.Print("remoteUser: ")
        fmt.Scan(&remoteUser)
        fmt.Print("remoteHost: ")
        fmt.Scan(&remoteHost)

	cmd := "cd go/src/github.com/Database;/usr/local/go/bin/go run main.go"
	out := ExecuteCommand(cmd)
	fmt.Println("output is", out)
	if strings.TrimSpace(out) == "No Flags Passed" {
		status = "healthy"
	} else {
		status = "unhealthy"
	}

	return status
}

func getKey() ssh.Signer {
	key, err := ioutil.ReadFile("./ec2.pem") //Make sure to rename this!!
	if err != nil {
		log.Fatalf("unable to read key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse key: %v", err)
	}

	return signer
}


// connects to ther machine
func connect() (*ssh.Client, *ssh.Session) {
	var port = "22"
	// get key
	signer := getKey()

	// configure authentication
	sshConfig := &ssh.ClientConfig{
		User: remoteUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// start a client connection to SSH server
	connection, err := ssh.Dial("tcp", remoteHost+":"+port, sshConfig)
	if err != nil {
		connection.Close()
		panic(err)
	}
	// create session
	session, err := connection.NewSession()
	if err != nil {
		session.Close()
		panic(err)
	}

	return connection, session
}

// ExecuteCommand runs commands passed to it in the other machine
func ExecuteCommand(cmd string) string {
	//connect to remote host
	connection, session := connect()
	// execute go program on remote host and get its combined standard output and standard error
	out, _ := session.CombinedOutput(cmd)
	defer connection.Close()
	defer session.Close()
	return string(out)
}
