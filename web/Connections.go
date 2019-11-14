package web

import (
	"fmt"
	"io/ioutil"
	"log"

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
	key, err := ioutil.ReadFile("./ec2.pem")
	if err != nil {
		log.Fatalf("unable to read key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse key: %v", err)
	}

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
	if string(out) == "Hello World" {
		status = "healthy"
	} else {
		status = "unhealthy"
	}

	defer connection.Close()
	defer session.Close()

	return status
}
