package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	store *Store
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {

	stations := a.store.All()

	writeJSON(w, http.StatusOK, stations)
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {

	country := r.PathValue("country")

	station, exists := a.store.Get(country)

	if !exists {
		writeError(w, http.StatusNotFound, "Station introuvable")
		return
	}

	writeJSON(w, http.StatusOK, station)
}

// createStation gère la route POST /stations
func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	var station Station

	err := json.NewDecoder(r.Body).Decode(&station)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalide")
		return
	}


	if a.store.Has(station.Country) {
		writeError(w, http.StatusConflict, "Cette station existe deja")
		return
	}

	a.store.Put(station)

	writeJSON(w, http.StatusCreated, station)
}