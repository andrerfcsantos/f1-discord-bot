package commands

import (
	"fmt"

	"f1-discord-bot/ergast"
)

// LastRace performs the actions for the "last" command sent to the bot,
// which informs the user about the results of the next grand prix.
// The result is a string ready to be sent to discord.
func LastRace() (string, error) {
	// Get next race from the API
	race, err := ergast.RequestLastRace()
	if err != nil {
		return "", fmt.Errorf("requesting last race to ergast: %v", err)
	}

	// Parse race time
	gpRFC3339Time := fmt.Sprintf("%sT%s", race.Date, race.Time)
	raceTime, err := ParseRFC3339InLocation(gpRFC3339Time, "Europe/Lisbon")
	if err != nil {
		return "", fmt.Errorf("parsing race time: %v", err)
	}

	// Build message
	var m TabularMessage

	m.Header = "Last Race results"
	m.Description = fmt.Sprintf("The last race was the %v at %v (%v, %v). The race was on %v.\nThe results are as follow:",
		race.RaceName,
		race.Circuit.CircuitName,
		race.Circuit.Location.Locality,
		race.Circuit.Location.Country,
		raceTime.Format("Monday, 02 January 2006 15:04 MST"))

	m.SetTableHeader("Pos", "Driver", "Constructor", "Time", "Fastest Lap", "Started")

	for _, result := range race.Results {
		m.AddRow(result.PositionText,
			result.Driver.FullName(),
			result.Constructor.Name,
			result.Time.Time,
			result.FastestLap.Time.Time,
			result.Grid)
	}

	return m.String(), nil
}
