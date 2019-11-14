package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

var remoteUser, remoteHost, port = "ubuntu", "ip", "22"

func firstConnect() string {
	var status string

	if remoteUser == "" {
		fmt.Print("remoteUser: ")
		fmt.Scan(&remoteUser)
		fmt.Print("remoteHost: ")
		fmt.Scan(&remoteHost)
	}

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
		log.Fatal("Unable to dial host:", err0)
	}
	// create session
	session, err1 := connection.NewSession()
	if err1 != nil {
		session.Close()
		log.Fatal("Unable to connect to host:", err1)
	}

	out, _ := session.CombinedOutput("go run main.go")
	if string(out) == "No Flags Passed" {
		status = "healthy"
	} else {
		status = "unhealthy"
	}

	defer connection.Close()
	defer session.Close()

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

func genLogin(cmd string) string {
	var response string

	// configure authentication
	signer := getKey()
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
		log.Fatal("Unable to dial host:", err0)
	}
	// create session
	session, err1 := connection.NewSession()
	if err1 != nil {
		session.Close()
		log.Fatal("Unable to connect to host:", err1)
	}

	out, err2 := session.CombinedOutput(cmd)
	if err2 != nil {
		log.Fatal("Unable to combine output:", err2)
	}

	response = strings.TrimSpace(string(out))

	defer connection.Close()
	defer session.Close()
	return response
}
