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
}
