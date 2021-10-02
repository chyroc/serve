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

func CreateLocalListener() (net.Listener, int, error) {
	port := 5000 // rand 100 step
	var err error
	var ln net.Listener
	for port < 10000 {
		ln, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			if strings.Contains(err.Error(), "bind: address already in use") {
				port += rand.Intn(100)
				continue
			}
			return nil, 0, err
		}
		return ln, port, nil
	}
	return nil, 0, err
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

func getEnableHttpsHost() ([]string, error) {
	ip, err := getOutboundIP()
	if err != nil {
		return nil, err
	}
	return []string{
		"localhost",
		"127.0.0.1",
		ip.String(),
		"::1",
	}, nil
}
