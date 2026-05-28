package main

type Store struct {
	Stations []Station
}

func NewStore(filepath string) (*Store, error) {
	stations, err := LoadFromJSON(filepath)
	if err != nil {
		return nil, err
	}
	return &Store{Stations: stations}, nil
}
