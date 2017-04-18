package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
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
func ResolveName(ip net.IP) ([]string, error) {
	// Try to get Neighbor DNS Names
	names, err := net.LookupAddr(ip.String())
	if err != nil {
		return nil, err
	}
	return names, nil
}

// Typeof Print type
func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
