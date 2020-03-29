package main

import "fmt"

// Bridge : a hue bridge
type Bridge struct {
	Username string
	Host     string
}

// CreateURI : creates a string, uri??? url??
func (b Bridge) CreateURI() string {
	return fmt.Sprintf("http://%s/api/%s/", b.Host, b.Username)
}
