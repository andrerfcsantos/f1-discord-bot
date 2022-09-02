package ergast

import (
	"fmt"
	"time"
)

var defaultLocation *time.Location

func init() {
	var err error
	defaultLocation, err = time.LoadLocation("Europe/Lisbon")
	if err != nil {
		panic(fmt.Errorf("loading default location: %w", err))
	}
}

// MRReply is the top level object present replies from the ergast API
type MRReply struct {
	MRData `json:"MRData"`
}

// MRData is a top level reply object containing some metadata about the reply
type MRData struct {
	Xmlns string `json:"xmlns"`
	// Series should always be "f1"
	Series string `json:"series"`
	// URL of the original request
	URL string `json:"url"`
	// Number of results in the reply
	Limit string `json:"limit"`
	// Number of results skipped
	Offset string `json:"offset"`
	// Total of records matching the request
	Total            string           `json:"total"`
	RaceTable        RaceTable        `json:"RaceTable"`
	CircuitTable     CircuitTable     `json:"CircuitTable"`
	DriverTable      DriverTable      `json:"DriverTable"`
	ConstructorTable ConstructorTable `json:"ConstructorTable"`
	SeasonTable      SeasonTable      `json:"SeasonTable"`
}

// Season represents a f1 season
type Season struct {
	Year string `json:"season"`
	// Wikipedia URL for the season
	URL string `json:"url"`
}

// SeasonTable represents a list of seasons
type SeasonTable struct {
	Seasons []Season `json:"Seasons"`
}

// HasSeason checks if a season with a given id is present on the season table
func (st *SeasonTable) HasSeason(year string) bool {
	for _, season := range st.Seasons {
		if season.Year == year {
			return true
		}
	}

	return false
}

// CircuitTable represents information about a list of circuits
type CircuitTable struct {
	Circuits []Circuit `json:"Circuits"`
}

// HasCircuit checks if a circuit with a given id is present on the circuit table
func (ct *CircuitTable) HasCircuit(circuitID string) bool {
	for _, circuit := range ct.Circuits {
		if circuit.CircuitID == circuitID {
			return true
		}
	}

	return false
}

// DriverTable contains a list of drivers
type DriverTable struct {
	Drivers []Driver `json:"Drivers"`
}

// HasDriver checks if a circuit with a given id is present on the circuit table
func (dt *DriverTable) HasDriver(driverID string) bool {
	for _, driver := range dt.Drivers {
		if driver.DriverID == driverID {
			return true
		}
	}

	return false
}

// ConstructorTable represents a list of constructors
type ConstructorTable struct {
	Constructors []Constructor `json:"Constructors"`
}

// HasConstructor checks if a circuit with a given id is present on the circuit table
func (dt *ConstructorTable) HasConstructor(constructorID string) bool {
	for _, constructor := range dt.Constructors {
		if constructor.ConstructorID == constructorID {
			return true
		}
	}

	return false
}

// RaceTable represents a list of races
type RaceTable struct {
	Season string `json:"season"`
	Round  string `json:"round"`
	Races  []Race `json:"Races"`
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
	Season   string       `json:"season"`
	Round    string       `json:"round"`
	URL      string       `json:"url"`
	RaceName string       `json:"raceName"`
	Circuit  Circuit      `json:"Circuit"`
	Results  []RaceResult `json:"Results"`
	DateTime
	FirstPractice  *DateTime `json:"FirstPractice"`
	SecondPractice *DateTime `json:"SecondPractice"`
	ThirdPractice  *DateTime `json:"ThirdPractice"`
	Qualifying     *DateTime `json:"Qualifying"`
	Sprint         *DateTime `json:"Sprint"`
}

// RaceResult represents a race result
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

type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func (dt *DateTime) GoTime() (time.Time, error) {
	t, err := time.Parse(time.RFC3339, dt.Date+"T"+dt.Time)
	if err != nil {
		return time.Time{}, fmt.Errorf("parsing time %q as RFC3339: %w", dt.Date+"T"+dt.Time, err)
	}
	return t, nil
}

func (dt *DateTime) TimeInDefaultLocation() time.Time {
	t, _ := dt.GoTime()
	return t.In(defaultLocation)
}

func (dt *DateTime) TimeInLocation(location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, fmt.Errorf("loading location %q: %w", location, err)
	}

	t, err := dt.GoTime()
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}

// MillisTime represents a time in miliseconds
type MillisTime struct {
	Millis string `json:"millis"`
	Time   string `json:"time"`
}

// Driver contains information about a f1 driver
type Driver struct {
	DriverID        string `json:"driverId"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
	Code            string `json:"code,omitempty"`
	PermanentNumber string `json:"permanentNumber,omitempty"`
}

// FullName returns the full name of a driver
func (d *Driver) FullName() string {
	return d.GivenName + " " + d.FamilyName
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

// Time represents a time of a fastest lap
type Time struct {
	Time string `json:"time"`
}

// AverageSpeed contains information about an average sped result
type AverageSpeed struct {
	Units string `json:"units"`
	Speed string `json:"speed"`
}
