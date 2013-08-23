package spipe

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Pipe struct {
	Network string

	ListenHost string
	ListenPort uint

	RemoteHost string
	RemotePort uint

	SourcePort uint
}

func (pipe *Pipe) LocalAddr() string {
	return fmt.Sprintf("%s:%d", pipe.ListenHost, pipe.ListenPort)
}

func (pipe *Pipe) RemoteAddr() string {
	return fmt.Sprintf("%s:%d", pipe.RemoteHost, pipe.RemotePort)
}

func (pipe *Pipe) Start() {
	listener, err := net.Listen("tcp", pipe.LocalAddr())
	if err != nil {
		panic(fmt.Sprintf("Error listening: %s", err.Error()))
		return
	}

	log.Println(fmt.Sprintf("Listening on %s...", pipe.LocalAddr()))

	for {
		// Accept a new incoming connection
		inConn, err := listener.Accept()
		if err != nil {
			log.Panicf("Error accepting: %s", err.Error())
			return
		}
		defer inConn.Close()

		log.Printf("Incoming connection from %s", inConn.RemoteAddr())

		// Start simultaneous outgoing connection
		addr, _ := net.ResolveTCPAddr("tcp", pipe.RemoteAddr())
		outConn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			log.Panicf("Connection to %s failed: %s", pipe.RemoteAddr(), err)
			return
		}
		defer outConn.Close()

		log.Printf("Piping %s <-> %s", inConn.RemoteAddr(), pipe.RemoteAddr())

		go pipeConnection(inConn, outConn)
		go pipeConnection(outConn, inConn)
	}
}

func pipeConnection(a net.Conn, b net.Conn) {
	defer a.Close()
	io.Copy(b, a)
}
