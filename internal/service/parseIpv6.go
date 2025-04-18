package service

import (
	"net"
)

// ParseIPv6 из [::1]:58635 в ::1.
func ParseIPv6(remoteAddr string) (string, error) {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return "", err
	}
	parsedIp := net.ParseIP(host)
	return parsedIp.String(), nil
}
