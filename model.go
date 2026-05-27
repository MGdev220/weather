package weather

type Station struct {
	country     string
	location    Location
	deviceModel string
	altitudeM   int
	Observation []Observation
}

type Location struct {
	longitude float64
	latitude  float64
}

type Observation struct {
	temperature  float64
	wind         Wind
	skycondition string
	notes        string
}
type Wind struct {
	Speed float64
	Deg   int
}
