package util

import (
	"net"
	"strings"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "N/A"
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}

		ip := ipNet.IP.To4()
		if ip == nil {
			continue
		}

		if strings.HasPrefix(ip.String(), "192.168.") {
			return ip.String()
		}
	}

	return "N/A"
}
