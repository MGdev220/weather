package main

import (
	"encoding/json"
	_ "encoding/json"
	"os"
	_ "os"
	_ "time"
)

var countryToISO = map[string]string{
	"France":    "FR",
	"Espagne":   "ES",
	"Portugal":  "PT",
	"Italie":    "IT",
	"Allemagne": "DE",
	"Belgique":  "BE",
	"Pays-Bas":  "NL",
	"Autriche":  "AT",
	"Suisse":    "CH",
	"Danemark":  "DK",
	"Suède":     "SE",
	"Norvège":   "NO",
	"Pologne":   "PL",
	"Tchéquie":  "CZ",
}

type jsonRoot struct {
	Stations []jsonStation `json:"stations"`
}

type jsonStation struct {
	ID           string            `json:"id"`
	Country      string            `json:"country"`
	Altitude     int               `json:"altitude_m"`
	Location     jsonLocation      `json:"location"`
	Device       jsonDevice        `json:"device"`
	Observations []jsonObservation `json:"observations"`
}

type jsonLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type jsonDevice struct {
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
}

type jsonObservation struct {
	Temperature float64  `json:"temperature_celsius"`
	Wind        jsonWind `json:"wind"`
	Conditions  string   `json:"conditions"`
	Notes       *string  `json:"notes"`
}

type jsonWind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction int     `json:"direction_deg"`
}

func (js *jsonStation) toModel() Station {
	// Conversion du pays
	isoCode, exists := countryToISO[js.Country]
	if !exists {
		isoCode = js.Country
	}

	var obsList []Observation
	for _, jo := range js.Observations {

		obsList = append(obsList, Observation{
			// Timestamp:    timestamp,
			Temperature: jo.Temperature,
			Condition:   jo.Conditions,
			Wind: Wind{
				Speed: jo.Wind.Speed,
				Deg:   jo.Wind.Direction,
			},
			Notes: jo.Notes,
		})
	}

	return Station{
		ID:      js.ID,
		Country: isoCode,
		Location: Location{
			Latitude:  js.Location.Latitude,
			Longitude: js.Location.Longitude,
		},
		DeviceModel: js.Device.Type,
		AltitudeM:   js.Altitude,
		Observation: obsList,
	}
}

func LoadFromJSON(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root jsonRoot
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	var stations []Station
	for _, js := range root.Stations {
		stations = append(
			stations,
			js.toModel(),
		)
	}

	return stations, nil
}
