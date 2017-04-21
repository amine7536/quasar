package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// IsValidPath - Check if a given path is valid
func IsValidPath(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}

// ResolveName get neighbor ip
func ResolveName(ip string) ([]string, error) {
	// Try to get Neighbor DNS Names
	names, err := net.LookupAddr(ip)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// ResolveNilrName when netmask is "/32""
func ResolveNilrName(network string) ([]string, error) {
	t := strings.Split(network, "/")
	if t[1] == "32" {
		return ResolveName(t[0])
	}
	return nil, errors.New("Not a /32 netmask")
}

// Typeof Print type
func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
