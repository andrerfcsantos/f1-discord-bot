package commands

import (
	"f1-discord-bot/ergast"
	"fmt"
	"time"
)

// NextRace performs the actions for the "next" command sent to the bot,
// which informs the user about the next grand prix. The result is a string ready to
// be sent to discord.
func NextRace() (string, error) {

	// Get next race from the API
	race, err := ergast.RequestNextRace()
	if err != nil {
		return "", fmt.Errorf("requesting next race to ergast: %v", err)
	}

	// Parse race time
	gpRFC3339Time := fmt.Sprintf("%sT%s", race.Date, race.Time)
	raceTime, err := ParseRFC3339InLocation(gpRFC3339Time, "Europe/Lisbon")
	if err != nil {
		return "", fmt.Errorf("parsing race time: %v", err)
	}

	// Build message
	var m HeaderMessage

	m.Header = "Next Race Information"
	m.Description = fmt.Sprintf("The next race is the %v at %v (%v, %v). The race will be on %v (%v). %v",
		race.RaceName,
		race.Circuit.CircuitName,
		race.Circuit.Location.Locality,
		race.Circuit.Location.Country,
		raceTime.Format("Monday, 02 January 2006 15:04 MST"),
		PrettyCountdount(time.Until(raceTime)),
		RaceHourComment(raceTime))

	return m.String(), nil
}
