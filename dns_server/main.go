package main

import "net"

func main() {
	serve()
}

func serve() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 53,
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
