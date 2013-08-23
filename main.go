package main

import (
	"flag"
	"fmt"
	spipe "github.com/s1kx/spipe/lib"
	"os"
	"strconv"
)

func main() {
	// Start options
	var (
		listenHost = flag.String("i", "0.0.0.0", "Listening interface IP address")
		sourcePort = flag.Uint("s", 0, "Outbound source port number")
	)
	flag.Parse()

	if flag.NArg() < 3 {
		fmt.Println("Usage: spipe [options] <listenPort> <remoteHost> <remotePort>")
		fmt.Println("")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Example: spipe -i 0.0.0.0 80 example.com 80")
		return
	}

	listenPort, err := strconv.Atoi(flag.Arg(0))
	if err != nil || listenPort <= 0 || listenPort >= 65536 {
		fmt.Println(fmt.Sprintf("Invalid listening port: %s", flag.Arg(0)))
		os.Exit(-1)
	}
	remoteHost := flag.Arg(1)
	remotePort, err := strconv.Atoi(flag.Arg(2))
	if err != nil || remotePort <= 0 || remotePort >= 65536 {
		fmt.Println(fmt.Sprintf("Invalid remote port: %s", flag.Arg(0)))
		os.Exit(-1)
	}

	pipe := &spipe.Pipe{*listenHost, uint(listenPort), remoteHost, uint(remotePort), *sourcePort}
	pipe.Start()
}
