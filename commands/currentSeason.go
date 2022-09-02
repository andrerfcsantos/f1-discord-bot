package commands

import (
	"fmt"

	"f1-discord-bot/ergast"
)

// CurrentSeason builds the message for the "current" command
func CurrentSeason() (string, error) {
	// Get next race from the API
	rt, err := ergast.CurrentSeason()
	if err != nil {
		return "", fmt.Errorf("requesting last race to ergast: %v", err)
	}

	// Buld message
	var m TabularMessage

	m.Header = "Races for the current season"
	m.SetTableHeader("Round", "Circuit", "Location", "Country", "Time")

	for _, race := range rt.Races {
		var localTimeStr string
		// Parse race time
		gpRFC3339Time := fmt.Sprintf("%sT%s", race.Date, race.Time)

		t, err := ParseRFC3339InLocation(gpRFC3339Time, "Europe/Lisbon")
		if err != nil {
			localTimeStr = gpRFC3339Time
		} else {
			localTimeStr = t.Format("02 Jan 15:04 MST")
		}

		m.AddRow(race.Round,
			race.Circuit.CircuitName,
			race.Circuit.Location.Locality,
			race.Circuit.Location.Country,
			localTimeStr)
	}

	return m.String(), nil
}
