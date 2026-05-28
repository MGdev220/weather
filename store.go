package main

type Store struct {
	stations map[string]Station
}

func NewStore() *Store {
	return &Store{
		stations: make(map[string]Station),
	}
}

func (s *Store) Put(st Station) {
	s.stations[st.ID] = st
}

func (s *Store) Has(id string) bool {
	_, exists := s.stations[id]
	return exists
}

func (s *Store) Get(id string) (Station, bool) {
	station, exists := s.stations[id]
	return station, exists
}

func (s *Store) Delete(id string) bool {
	if s.Has(id) {
		delete(s.stations, id)
		return true
	}
	return false
}

func (s *Store) All() []Station {
	var all []Station
	for _, st := range s.stations {
		all = append(all, st)
	}
	return all
}
