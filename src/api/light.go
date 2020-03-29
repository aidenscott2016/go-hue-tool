package api

// Light : a light
type Light struct {
	ID    int
	State LightState
}

// LightState : attributes for a light's state
type LightState struct {
	On bool `json:"on"`
}

func (l Light) updateOn(on bool) {
	l.State.On = on
}
