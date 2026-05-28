package main

import (
	"fmt"
	"log"
)

func main() {

	stations, err := LoadFromJSON("weather_data.json")
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier JSON : %v", err)
	}

	fmt.Printf("Nombre total de stations : %d\n", len(stations))

	if len(stations) > 0 {
		fmt.Printf("Nombre d'observations pour la 1ere station : %d\n", len(stations[0].Observation))

	} else {
		fmt.Println("\n la slice est vide")
	}

	xmlStations, err := LoadFromXML("weather_data.xml")
	if err != nil {
		log.Fatalf("Erreur XML : %v", err)
	}

	fmt.Printf("Total stations JSON : %d | XML : %d\n", len(stations), len(xmlStations))

	if len(stations) == len(xmlStations) && len(stations) > 0 {

		obsJson := len(stations[0].Observation)
		obsXml := len(xmlStations[0].Observation)
		fmt.Printf("Total observations JSON : %d | XML : %d\n", obsJson, obsXml)

		// Comparaison de la première température
		if obsJson > 0 && obsXml > 0 {
			tempJson := stations[0].Observation[0].Temperature
			tempXml := xmlStations[0].Observation[0].Temperature
			fmt.Printf("Température JSON : %.2f | XML : %.2f\n", tempJson, tempXml)
		}

	} else {
		fmt.Println("Erreur : Le nombre de stations est différent entre JSON et XML.")
	}
}
