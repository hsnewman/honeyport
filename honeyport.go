//
// Honeyport server
// Create a server which monitors tcp ports for connection
// If there is a connection event, log the event (but do nothing on the connection)
// command line: honeyport connection type portlow porthigh
//

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	CONN_HOST = ""
)

func main() {
	//
	// Get command line input
	//
	//	conn_type := os.Args[1]
	//	cpla := os.Args[2]
	//	cpha := os.Args[3]

	//	conn_port_low, _ := strconv.Atoi(cpla)
	//	conn_port_high, _ := strconv.Atoi(cpha)

	//
	//for testing
	//
	conn_port_low := 55515
	conn_port_high := 55530
	conn_type := "tcp"
	//
	// Spawn a thread for each port
	//   Loop through all ports to open them
	//   If someone connects, report it
	//
	for i := conn_port_low; i <= conn_port_high; i++ {
		fmt.Println("looping to handle port:", i)
		go handlePort(conn_type, i)
	}
	//
	// Now just sleep
	//
	for {
		time.Sleep(5000)
	}
}

//
// Open the port and service it
//
func handlePort(conn_type string, port int) {
	fmt.Println("in handleport with port:", port)
	l, err := net.Listen(conn_type, ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + strconv.Itoa(port))

	// Listen for an incoming connection.
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error(), port)
			os.Exit(1)
		}
		//logs an connection
		fmt.Printf("Connection from: %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
	}
}
