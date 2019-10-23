package discovery

import (
	"github.com/cakturk/go-netstat/netstat"
)

// ToDo: add functions for IPv6 connections

// GetTCPConnections returns a slice describing TCP connections
func GetTCPConnections() ([]netstat.SockTabEntry, error) {

	// TCP sockets
	socks, err := netstat.TCPSocks(netstat.NoopFilter)
	if err != nil {
		return nil, err
	}
	return socks, err
}

// GetUDPConnections returns a slice describing UDP connections
func GetUDPConnections() ([]netstat.SockTabEntry, error) {
	// UDP sockets
	socks, err := netstat.UDPSocks(netstat.NoopFilter)
	if err != nil {
		return nil, err
	}
	return socks, err
}
