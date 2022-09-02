package commands

import (
	"fmt"

	"f1-discord-bot/ergast"
)

// Results performs the actions for the "results" command sent to the bot
func Results(args ...string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("command 'results' needs more arguments")
	}

	subCommand := args[0]

	switch subCommand {
	case "circuit":
		switch {
		case len(args) < 2:
			return "", fmt.Errorf("command 'results circuit' needs a circuitID as an argument")
		case len(args) == 2:
			message, err := CircuitResults(args[1], 10)
			if err != nil {
				return "", fmt.Errorf("getting circuit results: %v", err)
			}
			return message, nil
		default:
			return "", fmt.Errorf("invalid number of arguments for the command 'results'")
		}

	case "driver":
		switch {
		case len(args) < 2:
			return "", fmt.Errorf("command 'results driver' needs a driverID as an argument")
		case len(args) == 2:
			message, err := DriverResults(args[1], 10)
			if err != nil {
				return "", fmt.Errorf("getting driver results: %v", err)
			}
			return message, nil
		default:
			return "", fmt.Errorf("invalid number of arguments for the command 'results'")
		}
	default:
		return "", fmt.Errorf("subcommand '%s' of 'results' not recognized or not yet implemented", subCommand)
	}
}

// CircuitResults performs the actions for the "results circuit <circuitID>" command sent to the bot
func CircuitResults(circuitID string, n int) (string, error) {
	// Get circuits
	circuitTable, err := ergast.Circuits()
	if err != nil {
		return "", fmt.Errorf("getting list of circuits from ergast: %v", err)
	}

	if !circuitTable.HasCircuit(circuitID) {
		// The circuit requested was not found in the list of circuits
		// Compute the levenshtein distances between the given argument and
		// all the circuits to see the one the user probably meant
		var lds LevenshteinDistances

		for _, circuit := range circuitTable.Circuits {
			lds = append(lds, LevenshteinDistance{Str1: circuitID, Str2: circuit.CircuitID})
		}
		lds.ComputeAll()
		lds.SortByDistance()
		return fmt.Sprintf("**UPS!**\nNo circuit with id '%s' was found.\nDid you mean?\n\t- %s", circuitID, lds[0].Str2), nil
	}

	// Get circuit results from the API
	raceTable, err := ergast.RequestCircuitResults(circuitID)
	if err != nil {
		return "", fmt.Errorf("requesting circuit results to ergast: %v", err)
	}

	// Trim the first races
	nRaces := len(raceTable.Races)
	if nRaces < n {
		n = nRaces
	}
	races := raceTable.Races[(nRaces - n):]
	circuit := races[0].Circuit

	// Buld message
	var m TabularMessage

	m.Header = fmt.Sprintf("WINNERS IN THE LAST %d RACES AT %s", n, circuit.CircuitName)
	m.SetTableHeader("Year", "Driver", "Constructor", "Time", "Laps")

	for i := len(races) - 1; i >= 0; i-- {
		race := races[i]
		m.AddRow(race.Season,
			race.Results[0].Driver.FullName(),
			race.Results[0].Constructor.Name,
			race.Results[0].Time.Time,
			race.Results[0].Laps)
	}

	return m.String(), nil
}

// DriverResults performs the actions for the "results driver <driverID>" command sent to the bot
func DriverResults(driverID string, n int) (string, error) {
	// Get circuits
	driverTable, err := ergast.Drivers()
	if err != nil {
		return "", fmt.Errorf("getting list of circuits from ergast: %v", err)
	}

	if !driverTable.HasDriver(driverID) {
		var lds LevenshteinDistances

		for _, driver := range driverTable.Drivers {
			lds = append(lds, LevenshteinDistance{Str1: driverID, Str2: driver.DriverID})
		}
		lds.ComputeAll()
		lds.SortByDistance()
		return fmt.Sprintf("**UPS!**\nNo driver with id '%s' was found.\nDid you mean?\n\t- %s", driverID, lds[0].Str2), nil
	}

	// Get driver results from the API
	raceTable, err := ergast.RequestDriverResults(driverID)
	if err != nil {
		return "", fmt.Errorf("requesting circuit results to ergast: %v", err)
	}

	// Trim the first races
	nRaces := len(raceTable.Races)
	if nRaces < n {
		n = nRaces
	}
	races := raceTable.Races[(nRaces - n):]
	driver := races[0].Results[0].Driver

	// Build message
	var m TabularMessage

	m.Header = fmt.Sprintf("LAST %d RACE RESULTS FOR %s", n, driver.FullName())
	m.SetTableHeader("Year", "GP", "Pos.", "Grid", "Constructor", "Time (ms)", "Laps", "Status")

	for i := len(races) - 1; i >= 0; i-- {
		race := races[i]
		m.AddRow(race.Season,
			race.RaceName,
			race.Results[0].Position,
			race.Results[0].Grid,
			race.Results[0].Constructor.Name,
			race.Results[0].Time.Millis,
			race.Results[0].Laps,
			race.Results[0].Status)
	}

	return m.String(), nil
}
