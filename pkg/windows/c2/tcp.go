package c2

import (
	"bufio"
	"net"
)

// TCPConnect connects to a remote TCP server using the format: "127.0.0.1:8080"
func TCPConnect(ipAddr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", ipAddr)
	if err != nil {
		return nil, err
	}
	return conn, err
}

// TCPWrite writes arbitrary data to a socket
func TCPWrite(conn net.Conn, data []byte) error {
	// send data to server
	_, err := conn.Write(data)
	if err != nil {
		return err
	}
	return err
}

// TCPRead reads arbitrary data from a socket, using a newline as a delimiter
func TCPRead(conn net.Conn) ([]byte, error) {
	// read server response until newline
	response, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	return response, err
}
