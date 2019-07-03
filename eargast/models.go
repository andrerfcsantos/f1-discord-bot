package eargast

// MRReply is the top level object present replies from the eargast API
type MRReply struct {
	MRData MRData `json:"MRData"`
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

// Race represents a grand prix event in f1
type Race struct {
	Season   string  `json:"season"`
	Round    string  `json:"round"`
	URL      string  `json:"url"`
	RaceName string  `json:"raceName"`
	Circuit  Circuit `json:"Circuit"`
	Date     string  `json:"date"`
	Time     string  `json:"time"`
}

// RaceTable represents a list of races
type RaceTable struct {
	Season string `json:"season"`
	Round  string `json:"round"`
	Races  []Race `json:"Races"`
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
