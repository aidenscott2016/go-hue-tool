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
	var opts = getopt.CommandLine
	cmd := ""
	opts.Parse(os.Args)
	if opts.NArgs() > 0 {
		cmd = opts.Arg(0)
		opts.Parse(opts.Args())
	}

	if "" == *bridgeIP {
		fmt.Println("bridge is required")
		os.Exit(1)
	}
	if "" == *username {
		fmt.Println("username is required")
		os.Exit(1)
	}

	if -1 == *lamp {
		fmt.Println("lamp is required")
		os.Exit(1)
	}

	bridge := Bridge{Username: *username, Host: *bridgeIP}
	light := NewLight(*lamp, bridge)

	switch cmd {
	case "off":
		light.TurnOff()
		break
	case "on":
		light.TurnOn()
		break
	case "step":
		if 0 == *step {
			fmt.Println("increment is required")
			os.Exit(1)
		}
		light.IncrementBrightness(*step)
		break
	default:
		fmt.Println("cmd of off|on|step required")
	}

}
