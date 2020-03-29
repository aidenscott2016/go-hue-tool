package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Light : a Light
type Light struct {
	ID     int
	Bridge Bridge
	State  LightState
}

// LightState : attributes for a light's state
type LightState struct {
	On bool `json:"on"`
}

// NewLight : returns a light of ID
func NewLight(id int, bridge Bridge) Light {
	return Light{ID: id, Bridge: bridge, State: LightState{}}
}

// TurnOff : sets light state off
func (l Light) TurnOff() {
	l.State.On = false

}

// TurnOn : sets light state on
func (l Light) TurnOn() {
	l.State.On = true
}

func put(l Light) (bool, error) {
	client := &http.Client{}

	body, err := json.Marshal(l.State)
	if err != nil {
		return false, err
	}
	request, err := http.NewRequest("PUT", l.getStatusEndpoint(), bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	res, err := client.Do(request)
	if err != nil {
		return false, err
	}
	io.Copy(os.Stdout, res.Body)

	return true, nil
}

func (l Light) getStatusEndpoint() string {
	return fmt.Sprintf(l.Bridge.CreateURI()+"/%i/status", l.ID)
}
