package main

import (
	"fmt"
	"os"
	"net"
	"code.google.com/p/go.crypto/ssh"
)

const (
	authsock = "SSH_AUTH_SOCK"
)

func main() {
	s := os.Getenv(authsock)
	raddr,err := net.ResolveUnixAddr("unix", s)
	if err != nil {
		fmt.Print(err)
		return
	}

	c,err := net.DialUnix("unix", nil, raddr)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer c.Close()

	a := ssh.NewAgentClient(c)
	idents,err := a.RequestIdentities()
	if err != nil {
		fmt.Print(err)
		return
	}
	for _,i := range idents {
		fmt.Println(i.Comment)
	}
}
