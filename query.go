package main

func FilterByCountry(stations []Station, iso string) []Station {
	var result []Station
	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}
	return result
}

func AvgTemperature(s Station) float64 {
	if len(s.Observation) == 0 {
		return 0.0
	}

	var total float64
	for _, obs := range s.Observation {
		total += obs.Temperature
	}

	return total / float64(len(s.Observation))
}

func MaxWindGust(stations []Station) (Station, float64) {
	var maxGust float64
	var windiestStation Station

	for _, s := range stations {
		for _, obs := range s.Observation {
			if obs.Wind.Speed > maxGust {
				maxGust = obs.Wind.Speed
				windiestStation = s
			}
		}
	}

	return windiestStation, maxGust
}

func CountByCountry(stations []Station) map[string]int {
	counts := make(map[string]int)
	for _, s := range stations {
		counts[s.Country]++
	}
	return counts
}
