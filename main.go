package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	mode := flag.String("mode", "call", "use [call|listen] for change mode of helper")
	host := flag.String("host", "127.0.0.1", "set up host, which we will call for checking is it ready yet")
	port := flag.String("port", "60911", "set up port, which we will call for checking is it ready yet, or listen in \"listen\" mode")

	flag.Parse()

	modeString := *mode
	hostString := *host
	portString := *port

	if modeString == "call" {
		for {
			conn, _ := net.Dial("tcp", net.JoinHostPort(hostString, portString))
			if conn != nil {
				conn.Close()
				break
			}
		}
	} else {
		ln, err := net.Listen("tcp", "127.0.0.1:"+portString)
		if err == nil {
			for {
				ln.Accept()
			}
		} else {
			fmt.Print(err)
		}
	}
}
