package main

type Station struct {
	ID          string        `json:"id"`
	Country     string        `json:"country"`
	Location    Location      `json:"location"`
	DeviceModel string        `json:"device_model"`
	AltitudeM   int           `json:"altitude_m"`
	Observation []Observation `json:"observation"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Altitude  int
}

type Observation struct {
	Temperature float64 `json:"temperature"`
	Wind        Wind    `json:"wind"`
	Condition   string  `json:"condition"`
	Notes       *string `json:"notes"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}
