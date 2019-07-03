package eargast

// MRReply is the top level object present replies from the eargast API
type MRReply struct {
	MRData MRData `json:"MRData"`
}

// MRData is a top level reply object containing some metadata about the reply
type MRData struct {
	Xmlns     string    `json:"xmlns"`
	Series    string    `json:"series"`
	URL       string    `json:"url"`
	Limit     string    `json:"limit"`
	Offset    string    `json:"offset"`
	Total     string    `json:"total"`
	RaceTable RaceTable `json:"RaceTable"`
}

// Location represents the location of a grand prix
type Location struct {
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Locality string `json:"locality"`
	Country  string `json:"country"`
}

// Circuit represents a circuit where a grand prix can be held
type Circuit struct {
	CircuitID   string   `json:"circuitId"`
	URL         string   `json:"url"`
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"Location"`
}

// RaceTable represents a list of races
type RaceTable struct {
	Season string `json:"season"`
	Round  string `json:"round"`
	Races  []Race `json:"Races"`
}

// Race represents a grand prix event in f1
type Race struct {
	Season   string       `json:"season"`
	Round    string       `json:"round"`
	URL      string       `json:"url"`
	RaceName string       `json:"raceName"`
	Circuit  Circuit      `json:"Circuit"`
	Date     string       `json:"date"`
	Time     string       `json:"time"`
	Results  []RaceResult `json:"Results"`
}

type RaceResult struct {
	Number       string      `json:"number"`
	Position     string      `json:"position"`
	PositionText string      `json:"positionText"`
	Points       string      `json:"points"`
	Driver       Driver      `json:"Driver"`
	Constructor  Constructor `json:"Constructor"`
	Grid         string      `json:"grid"`
	Laps         string      `json:"laps"`
	Status       string      `json:"status"`
	Time         MillisTime  `json:"Time,omitempty"`
	FastestLap   FastestLap  `json:"FastestLap"`
}

type MillisTime struct {
	Millis string `json:"millis"`
	Time   string `json:"time"`
}

// Driver represents a f1 driver
type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}

// Constructor represents a f1 constructor
type Constructor struct {
	ConstructorID string `json:"constructorId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Nationality   string `json:"nationality"`
}

// FastestLap represents a fastest lap result from a driver in a race
type FastestLap struct {
	Rank         string       `json:"rank"`
	Lap          string       `json:"lap"`
	Time         Time         `json:"Time"`
	AverageSpeed AverageSpeed `json:"AverageSpeed"`
}

// Time
type Time struct {
	Time string `json:"time"`
}

type AverageSpeed struct {
	Units string `json:"units"`
	Speed string `json:"speed"`
}
