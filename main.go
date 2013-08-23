package main

import (
	"flag"
	"fmt"
	spipe "github.com/s1kx/spipe/lib"
	"log"
	"os"
	"strconv"
)

func main() {
	// Start options
	var (
		listenHost = flag.String("i", "0.0.0.0", "Listening interface IP address")
		sourcePort = flag.Uint("s", 0, "Outbound source port number")
		network    = flag.String("net", "tcp", "Network type (tcp, tcp4, tcp6)")
	)
	flag.Parse()

	// Check if all mandatory arguments were supplied
	if flag.NArg() < 3 {
		fmt.Println("spipe [options] <listen port> <remote host> <remote port>")
		fmt.Println("")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Example: spipe -i 127.0.0.1 80 example.com 80")
		os.Exit(1)
	}

	// Read mandatory arguments
	listenPort, err := strconv.Atoi(flag.Arg(0))
	if err != nil || listenPort < 1 || listenPort > 65535 {
		log.Fatalf("Invalid listening port: %s\n", flag.Arg(0))
	}
	remoteHost := flag.Arg(1)
	remotePort, err := strconv.Atoi(flag.Arg(2))
	if err != nil || remotePort < 1 || remotePort > 65535 {
		log.Fatalf("Invalid remote port: %s\n", flag.Arg(2))
	}

	// Start listener and pipe connections
	pipe := &spipe.Pipe{*network, *listenHost, uint(listenPort), remoteHost, uint(remotePort), *sourcePort}
	pipe.Start()
}
