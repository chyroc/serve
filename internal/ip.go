package internal

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetAvailableAddress() (net.Listener, error) {
	ip, err := getOutboundIP()
	if err != nil {
		return nil, err
	}
	listener, err := getAvailablePort(ip.String())
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

func getAvailablePort(ip string) (net.Listener, error) {
	port := 5000 // rand 100 step
	var err error
	var ln net.Listener
	for port < 10000 {
		ln, err = net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
		if err != nil {
			if strings.Contains(err.Error(), "bind: address already in use") {
				port += rand.Intn(100)
				continue
			}
			return nil, err
		}
		return ln, nil
	}
	return nil, err
}
