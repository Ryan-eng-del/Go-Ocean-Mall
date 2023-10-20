package util

import "net"

func GenRandomPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")

	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	return port

}
