package main

import (
	"encoding/xml"
	"os"
	"strconv"
)

type rootStation struct {
	Stations []xmlStation `xml:"station"`
}

type xmlStation struct {
	Country      string           `xml:"country,attr"`
	Coordinates  xmlCoordinates   `xml:"coordinates"`
	Hardware     xmlHardware      `xml:"hardware"`
	Observations []xmlObservation `xml:"observations>observation"`
}

type xmlCoordinates struct {
	Latitude  float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
	Altitude  int     `xml:"altitude,attr"`
}
type xmlHardware struct {
	Model string `xml:"model"`
}
type xmlMeasure struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}
type xmlWind struct {
	Speed float64 `xml:",chardata"`
	Deg   int     `xml:",chardata"`
}

type xmlObservation struct {
	Sky      string       `xml:"sky,attr"`
	Measures []xmlMeasure `xml:"measure"`
	Wind     xmlWind      `xml:"wind"`
	Note     *string      `xml:"note"`
}

func (xs *xmlStation) toModel() Station {
	var obsList []Observation

	for _, xo := range xs.Observations {
		var temp float64

		for _, m := range xo.Measures {
			if m.Type == "temperature" {

				parsedTemp, err := strconv.ParseFloat(m.Value, 64)
				if err == nil {
					temp = parsedTemp
				}
			}
		}

		obsList = append(obsList, Observation{
			Temperature: temp,
			Condition:   xo.Sky,
			Wind: Wind{
				Speed: xo.Wind.Speed,
				Deg:   xo.Wind.Deg,
			},
			Notes: xo.Note,
		})
	}

	return Station{
		Country: xs.Country,
		Location: Location{
			Latitude:  xs.Coordinates.Latitude,
			Longitude: xs.Coordinates.Longitude,
			Altitude:  xs.Coordinates.Altitude,
		},
		DeviceModel: xs.Hardware.Model,
		Observation: obsList,
	}
}

func LoadFromXML(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root rootStation

	if err := xml.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	var stations []Station
	for _, xs := range root.Stations {
		stations = append(stations, xs.toModel())
	}

	return stations, nil
}
