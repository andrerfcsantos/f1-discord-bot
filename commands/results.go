package commands

import (
	"f1-discord-bot/eargast"
	"fmt"
	"sort"
	"strings"
	tbw "text/tabwriter"

	"github.com/agnivade/levenshtein"
)

type HammingDistance struct {
	Value    string
	Distance int
}

type HammingDistances []HammingDistance

func (ds HammingDistances) SortByDistance() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Distance < ds[j].Distance
	})
}

// Results performs the actions for the "results" command sent to the bot
func Results(args ...string) (string, error) {

	if len(args) < 1 {
		return "", fmt.Errorf("command 'results' needs more arguments")
	}

	subCommand := args[0]

	switch subCommand {
	case "circuit":

		if len(args) == 1 {
			return "", fmt.Errorf("command 'results circuit' needs a circuitID as an argument")
		}

		message, err := CircuitResults(args[1], 10)

		if err != nil {
			return "", fmt.Errorf("getting circuit results: %v", err)
		}

		return message, nil
	default:
		return "", fmt.Errorf("command subcommand '%s' of 'results' not recognized or not yet implemented", subCommand)

	}
}

// CircuitResults performs the actions for the "results circuit <circuitID>" command sent to the bot
func CircuitResults(circuitID string, n int) (string, error) {
	var validCircuit bool
	var message strings.Builder
	var distances HammingDistances

	mrdata, err := eargast.Circuits()
	if err != nil {
		return "", fmt.Errorf("getting list of circuits from eargast: %v", err)
	}

	if len(mrdata.MRData.CircuitTable.Circuits) == 0 {
		return "", fmt.Errorf("empty list of circuits from eargast: %v", err)
	}

	for _, circuit := range mrdata.MRData.CircuitTable.Circuits {
		if circuit.CircuitID == circuitID {
			validCircuit = true
			break
		}

		distance := levenshtein.ComputeDistance(circuitID, circuit.CircuitID)
		distances = append(distances, HammingDistance{Value: circuit.CircuitID, Distance: distance})
	}

	if !validCircuit {
		distances.SortByDistance()
		return fmt.Sprintf("No circuit with id '%s' was found.\nMaybe you meant?\n\t- %s", circuitID, distances[0].Value), nil
	}

	// Get next race from the API
	raceInfo, err := eargast.RequestCircuitResults(circuitID)
	if err != nil {
		return "", fmt.Errorf("requesting circuit results to eargast: %v", err)
	}

	races := raceInfo.MRData.RaceTable.Races
	nRaces := len(races)
	// Even though the reply from the API came out ok, check if it actually returned any race
	if nRaces < 1 {
		return "", fmt.Errorf("request to eargast ok, but no races returned")
	}

	if nRaces < n {
		n = nRaces
	}

	races = races[(nRaces - n):]
	c := races[0].Circuit
	message.WriteString(fmt.Sprintf("** WINNERS IN THE LAST %d RACES AT %s **\n", n, strings.ToUpper(c.CircuitName)))
	message.WriteString("```")
	tableWriter := tbw.NewWriter(&message, 0, 0, 3, ' ', 0)

	fmt.Fprintln(tableWriter, "Year\tDriver\tConstructor\tTime\tLaps")

	for i := len(races) - 1; i >= 0; i-- {
		race := races[i]
		fmt.Fprintf(tableWriter, "%s\t%s\t%s\t%s\t%s\n",
			race.Season,
			race.Results[0].Driver.GivenName+" "+race.Results[0].Driver.FamilyName,
			race.Results[0].Constructor.Name,
			race.Results[0].Time.Time,
			race.Results[0].Laps)
	}

	tableWriter.Flush()
	message.WriteString("```")

	return message.String(), nil
}
