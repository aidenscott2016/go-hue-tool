package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Light : a Light
type Light struct {
	ID     int
	Bridge Bridge
	State  LightState
}

// LightState : attributes for a light's state
type LightState struct {
	On                  bool `json:"on,omitempty"`
	BrightnessIncrement int  `json:"bri_inc,omitempty"`
}

// NewLight : returns a light of ID
func NewLight(id int, bridge Bridge) Light {
	return Light{ID: id, Bridge: bridge, State: LightState{}}
}

// TurnOff : sets light state off
func (l Light) TurnOff() {
	l.State.On = false
	_, err := put(l)
	if nil != err {
		fmt.Println(err)
	}
}

// TurnOn : sets light state on
func (l Light) TurnOn() {
	l.State.On = true
	_, err := put(l)
	if nil != err {
		fmt.Println(err)
	}
}

// IncrementBrightness : increments brigtness
func (l Light) IncrementBrightness(inc int) {
	l.State.BrightnessIncrement = inc
	_, err := put(l)
	if nil != err {
		fmt.Println(err)
	}
}

func put(l Light) (bool, error) {
	client := http.Client{}

	body, err := json.Marshal(l.State)
	if err != nil {
		return false, err
	}

	request, err := http.NewRequest("PUT", l.getStateEndpoint(), bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}

	_, err = client.Do(request)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (l Light) getStateEndpoint() string {
	return fmt.Sprintf(l.Bridge.CreateURI()+"/lights/%d/state", l.ID)
}
