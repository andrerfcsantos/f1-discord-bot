package commands

import (
	"f1-discord-bot/eargast"
	"fmt"
	"strings"
	"time"

	tbw "text/tabwriter"
)

// LastRace performs the actions for the "last" command sent to the bot,
// which informs the user about the results of the next grand prix.
// The result is a string ready to be sent to discord.
func LastRace() (string, error) {
	var message string
	// Get next race from the API
	raceInfo, err := eargast.RequestLastRace()
	if err != nil {
		return "", fmt.Errorf("requesting last race to eargast: %v", err)
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

	// Build the message
	message = fmt.Sprintf("**LAST RACE RESULTS**\nThe last race was the %v at %v (%v, %v). The race was on %v.\nThe results are as follow: \n\n",
		race.RaceName,
		race.Circuit.CircuitName,
		race.Circuit.Location.Locality,
		race.Circuit.Location.Country,
		gpLocalTime.Format("Monday, 02 January 2006 15:04 MST"))

	var tablebBuilder strings.Builder

	tableWriter := tbw.NewWriter(&tablebBuilder, 0, 0, 3, ' ', 0)

	fmt.Fprintln(tableWriter, "Pos.\tDriver\tConstructor\tTime\tFastest Lap\tStarted")

	for _, result := range race.Results {
		fmt.Fprintf(tableWriter, "%s\t%s\t%s\t%s\t%s\t%s\n",
			result.PositionText,
			result.Driver.GivenName+" "+result.Driver.FamilyName,
			result.Constructor.Name,
			result.Time.Time,
			result.FastestLap.Time.Time,
			result.Grid)
	}

	tableWriter.Flush()

	message = message + "```" + tablebBuilder.String() + "```"

	return message, nil
}
