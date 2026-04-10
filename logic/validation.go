package logic

import (
	"net"
	"strconv"

	"github.com/pterm/pterm"
)

func ValidatePort(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		pterm.Error.Printfln("Invalid Port: %s is not a number", port)
		return false
	}

	if p < 1 || p > 65535 {
		pterm.Error.Println("Invalid Port: Port must be between 1 and 65535")
		return false
	}

	return true
}

func ValidateIP(ip string) bool {
	if net.ParseIP(ip) == nil {
		pterm.Error.Printfln("Invalid IP address: %s", ip)
		return false
	}

	return true
}
