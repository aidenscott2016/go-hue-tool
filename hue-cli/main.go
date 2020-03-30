package main

import (
	"fmt"
	"os"

	"github.com/pborman/getopt/v2"
)

func main() {
	bridgeIP := getopt.StringLong("bridge", 'b', "", "The IP of the bridge")
	username := getopt.StringLong("username", 'u', "", "The IP of the bridge")
	lamp := getopt.IntLong("lamp", 'l', -1, "The ID of the lamp")
	step := getopt.IntLong("increment", 's', 0, "The the amount to increment by")

	getopt.ParseV2()

	if *bridgeIP == "" {
		fmt.Println("bridge is required")
		os.Exit(1)
	}
	if *username == "" {
		fmt.Println("username is required")
		os.Exit(1)
	}

	bridge := Bridge{Username: *username, Host: *bridgeIP}
	light := NewLight(*lamp, bridge)
	light.IncrementBrightness(*step)
}
