package main

import (
	"fmt"
	"log"
	"net/http"
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

		if obsJson > 0 && obsXml > 0 {
			tempJson := stations[0].Observation[0].Temperature
			tempXml := xmlStations[0].Observation[0].Temperature
			fmt.Printf("Température JSON : %.2f | XML : %.2f\n", tempJson, tempXml)
		}

	} else {
		fmt.Println("Erreur : Le nombre de stations est différent entre JSON et XML.")
	}

	countObs := func(stations []Station) int {
		total := 0
		for _, s := range stations {
			total += len(s.Observation)
		}
		return total
	}
	obsJson := countObs(stations)
	obsXml := countObs(xmlStations)

	fmt.Printf("JSON : %d stations, %d observations\n", len(stations), obsJson)
	fmt.Printf("XML  : %d stations, %d observations\n", len(xmlStations), obsXml)

	if len(stations) == len(xmlStations) && obsJson == obsXml {
		fmt.Println("Cohérence : OK\n")
	} else {
		fmt.Println("Cohérence : ECHEC\n")
	}

	windiestStation, maxGust := MaxWindGust(stations)

	fmt.Printf("Station la plus ventée : %s (%.1f km/h)\n", windiestStation.Country, maxGust)

	var bordeaux Station
	for _, s := range stations {
		if s.Country == "FR-BOR-001" {
			bordeaux = s
			break
		}
	}
	avgTemp := AvgTemperature(bordeaux)
	fmt.Printf("Temp. moyenne Bordeaux Mérignac : %.1f °C\n", avgTemp)

	counts := CountByCountry(stations)
	fmt.Printf("Stations par pays : %v\n", counts)

	store, err := NewStore("weather_data.json")
	if err != nil {
		log.Fatalf("Erreur au chargement des données JSON : %v", err)
	}

	_ = &Server{store: store}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.ListenAndServe(":8080", mux)

	log.Println("Serveur démarré sur http://localhost:8080...")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Erreur  du serveur : %v", err)
	}

}
