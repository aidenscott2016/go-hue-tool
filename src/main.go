package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hue-cli/src/api"
	"io"
	"net/http"
	"os"
)

type lightState struct {
	On bool `json:"on"`
}

func main() {
	args := os.Args
	bridgeIP := args[1]
	username := args[2]
	lamp := args[3]

	switchOff(bridgeIP, username, lamp)
}

func switchOff2() {
	id := 6
	light := api.Light{ID: 6, State: api.LightState{On: false}}

}

func switchOff(bridge string, username string, lamp string) {
	client := &http.Client{}
	off := lightState{On: true}
	offBody, err := json.Marshal(off)
	if err != nil {
		panic(err)
	}
	uri := createURI(bridge, username, "lights") + lamp + "/state"
	fmt.Println(uri)
	request, err := http.NewRequest("PUT", uri, bytes.NewBuffer(offBody))
	if err != nil {
		panic(err)
	}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, res.Body)
}

func createURI(bridge string, username string, resource string) string {
	return fmt.Sprintf("http://%s/api/%s/%s/", bridge, username, resource)
}