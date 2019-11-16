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
	key, err0 := ioutil.ReadFile("./ec2.pem") 
	if err0 != nil {
		log.Fatalf("unable to read key: %v", err0)
	}
	signer, err1 := ssh.ParsePrivateKey(key)
	if err1 != nil {
		log.Fatalf("unable to parse key: %v", err1)
	}

	return signer
}


// connects to ther machine
func connect() (*ssh.Client, *ssh.Session) {
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
	connection, err0 := ssh.Dial("tcp", remoteHost+":"+port, sshConfig)
	if err0 != nil {
		connection.Close()
		panic(err0)
	}
	// create session
	session, err1 := connection.NewSession()
	if err1 != nil {
		session.Close()
		panic(err1)
	}

	return connection, session
}

// ExecuteCommand runs commands passed to it in the other machine
func ExecuteCommand(cmd string) string {
	//connect to remote host
	connection, session := connect()
	// execute go program on remote host and get its combined standard output and standard error
	out, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatal("Unable to get combined output:", err)
	}
	defer connection.Close()
	defer session.Close()
	return string(out)
}
