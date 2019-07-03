package commands

import (
	"f1-discord-bot/eargast"
	"fmt"
	"time"
)

// NextRace performs the actions for the "next" command sent to the bot,
// which informs the user about the next grand prix. The result is a string ready to
// be sent to discord.
func NextRace() (string, error) {

	// Get next race from the API
	raceInfo, err := eargast.RequestNextRace()
	if err != nil {
		return "", fmt.Errorf("requesting next race to eargast: %v", err)
	}

	// Even though the reply from the API came out ok, check if it actually returned any race
	if len(raceInfo.MRData.RaceTable.Races) < 1 {
		return "", fmt.Errorf("request to eargast ok, but no races returned")
	}

	race := raceInfo.MRData.RaceTable.Races[0]

	// Parse race time
	gpRFC3339Time := fmt.Sprintf("%sT%s", race.Date, race.Time)
	gpTime, err := time.Parse(time.RFC3339, gpRFC3339Time)

	if err != nil {
		return "", fmt.Errorf("parsing race time '%s': %v", gpTime, err)
	}

	// To provide more context to the user, convert the race time to their local time.
	// TODO: This bit of code is very opinionated about what the user timezone is and the result
	// probably can be cached or set at
	europeLisbonLoc, err := time.LoadLocation("Europe/Lisbon")
	if err != nil {
		return "", fmt.Errorf("loading 'Europe/Lisbon' location: %v", err)
	}
	gpLocalTime := gpTime.In(europeLisbonLoc)

	// Check if the user will need to wake up early for the race :)
	hour := gpLocalTime.Hour()
	var hourComment string

	if hour > 4 && hour < 9 {
		hourComment = "Unfortunately it seems you'll have to wake up early if you want to watch the race :("
	} else {
		hourComment = "It seems a decent hour for the race. You won't have to wake up early!"
	}

	// Build the message
	message := fmt.Sprintf("The next race is the %v at %v (%v, %v). The race will be on %v (%v). %v",
		race.RaceName,
		race.Circuit.CircuitName,
		race.Circuit.Location.Locality,
		race.Circuit.Location.Country,
		gpLocalTime.Format("Monday, 02 January 2006 15:04 MST"),
		PrettyCountdount(time.Until(gpLocalTime)),
		hourComment)

	return message, nil
}

// PrettyCountdount sees a Duration as a countdown for an event to happen
// and transforms into a easily readible string representation.
// The results contains information about if the event it's in the past or it's still to come.
// If the event it's still to come, it will also contain information about the time left,
// expressed in minutes, hours, or days depending on the value of the duration.
func PrettyCountdount(d time.Duration) string {

	switch {
	case d < 0:
		return "already over"
	case d.Hours() < 1:
		return fmt.Sprintf("%.1f minutes to go", d.Minutes())
	case d.Hours() < 24:
		return fmt.Sprintf("%.1f hours to go", d.Hours())
	default:
		return fmt.Sprintf("%.1f days to go", d.Hours()/24)
	}

}
