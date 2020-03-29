package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args
	bridgeIP := args[1]
	username := args[2]
	lamp, _ := strconv.Atoi(args[3])
	bridge := Bridge{Username: username, Host: bridgeIP}

	switchOff(lamp, bridge)
}

func switchOff(lightID int, bridge Bridge) {
	light := NewLight(lightID, bridge)
	light.TurnOff()

}

// func switchOffOld(bridge string, username string, lamp string) {
// 	client := &http.Client{}
// 	off := lightState{On: true}
// 	offBody, err := json.Marshal(off)
// 	if err != nil {
// 		panic(err)
// 	}
// 	uri := createURI(bridge, username, "lights") + lamp + "/state"
// 	fmt.Println(uri)
// 	request, err := http.NewRequest("PUT", uri, bytes.NewBuffer(offBody))
// 	if err != nil {
// 		panic(err)
// 	}
// 	res, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	io.Copy(os.Stdout, res.Body)
// }

// func createURI(bridge string, username string, resource string) string {
// 	return fmt.Sprintf("http://%s/api/%s/%s/", bridge, username, resource)
// }
