package service

import (
	"net"
)

// ParseIPv just delete port.
func ParseIPv(remoteAddr string) (string, error) {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return "", err
	}
	parsedIp := net.ParseIP(host)
	return parsedIp.String(), nil
}
