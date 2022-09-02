package ergast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// RequestNextRace uses the ergast api to request information the next race
func RequestNextRace() (Race, error) {
	reply, err := APIGet("/current/next.json")
	if err != nil {
		return Race{}, err
	}

	if len(reply.MRData.RaceTable.Races) == 0 {
		return Race{}, fmt.Errorf("request ok, but no races returned")
	}
	return reply.MRData.RaceTable.Races[0], nil
}

// RequestLastRace requests information about the last race
func RequestLastRace() (Race, error) {
	reply, err := APIGet("/current/last/results.json")
	if err != nil {
		return Race{}, err
	}
	if len(reply.MRData.RaceTable.Races) == 0 {
		return Race{}, fmt.Errorf("request ok, but no races returned")
	}
	return reply.MRData.RaceTable.Races[0], nil
}

// RequestCircuitResults requests information about results on a given circuit in the last years
func RequestCircuitResults(circuitID string) (RaceTable, error) {
	endpoint := fmt.Sprintf("/circuits/%s/results/1.json?limit=1000", strings.ToLower(circuitID))
	reply, err := APIGet(endpoint)
	if err != nil {
		return RaceTable{}, err
	}
	if len(reply.MRData.RaceTable.Races) == 0 {
		return RaceTable{}, fmt.Errorf("request ok, but no races returned")
	}
	return reply.MRData.RaceTable, nil
}

// RequestDriverResults requests information about results for a given driver in the last races
func RequestDriverResults(driverID string) (RaceTable, error) {
	endpoint := fmt.Sprintf("/drivers/%s/results.json?limit=1000", strings.ToLower(driverID))
	reply, err := APIGet(endpoint)
	if err != nil {
		return RaceTable{}, err
	}
	if len(reply.MRData.RaceTable.Races) == 0 {
		return RaceTable{}, fmt.Errorf("request ok, but no races returned")
	}
	return reply.MRData.RaceTable, nil
}

// CurrentSeason requests information about races of the current season
func CurrentSeason() (RaceTable, error) {
	reply, err := APIGet("/current.json?limit=1000")
	if err != nil {
		return RaceTable{}, err
	}
	if len(reply.MRData.RaceTable.Races) == 0 {
		return RaceTable{}, fmt.Errorf("request ok, but no races returned")
	}
	return reply.MRData.RaceTable, nil
}

// Circuits requests a list of circuits
func Circuits() (CircuitTable, error) {
	reply, err := APIGet("/circuits.json?limit=1000")
	if err != nil {
		return CircuitTable{}, err
	}
	if len(reply.MRData.CircuitTable.Circuits) == 0 {
		return CircuitTable{}, fmt.Errorf("empty list of circuits from ergast: %v", err)
	}
	return reply.MRData.CircuitTable, nil
}

// Drivers requests a list of drivers
func Drivers() (DriverTable, error) {
	reply, err := APIGet("/drivers.json?limit=1000")
	if err != nil {
		return DriverTable{}, err
	}
	if len(reply.MRData.DriverTable.Drivers) == 0 {
		return DriverTable{}, fmt.Errorf("empty list of drivers from ergast: %v", err)
	}
	return reply.MRData.DriverTable, nil
}

// Constructors requests a list of constructors
func Constructors() (ConstructorTable, error) {
	reply, err := APIGet("/constructors.json?limit=1000")
	if err != nil {
		return ConstructorTable{}, err
	}
	if len(reply.MRData.ConstructorTable.Constructors) == 0 {
		return ConstructorTable{}, fmt.Errorf("empty list of constructors from ergast: %v", err)
	}
	return reply.MRData.ConstructorTable, nil
}

// Seasons requests a list of seasons
func Seasons() (SeasonTable, error) {
	reply, err := APIGet("/seasons.json?limit=1000")
	if err != nil {
		return SeasonTable{}, err
	}
	if len(reply.MRData.SeasonTable.Seasons) == 0 {
		return SeasonTable{}, fmt.Errorf("empty list of seasons from ergast: %v", err)
	}
	return reply.MRData.SeasonTable, nil
}

// APIGet makes a GET request to the specified API endpoint.
func APIGet(endpoint string) (MRReply, error) {
	// Make the request
	reply, err := Client.Get(BaseURL + endpoint)
	if err != nil {
		return MRReply{}, fmt.Errorf("GET Request: %v", err)
	}

	// Read the reply bytes
	replyBytes, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return MRReply{}, fmt.Errorf("reading reply bytes: %v", err)
	}

	// Unmarshal the bytes into a MRReply
	var res MRReply
	err = json.Unmarshal(replyBytes, &res)
	if err != nil {
		return MRReply{}, fmt.Errorf("unmarshaling reply bytes into json: %v", err)
	}

	return res, nil
}
