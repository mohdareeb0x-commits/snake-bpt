package logic

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/exec"
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

func CheckDependency(name string) error {
	_, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("Dependency Error: %s is not installed or not in PATH", name)
	}

	return nil
}

func CheckInternet() error {
	resp, err := http.Get("https://clients3.google.com/generate_204")
	if err == nil {
		resp.Body.Close()
	}
	if err != nil {
		return errors.New("No internet connection available")
	}

	return nil
}
