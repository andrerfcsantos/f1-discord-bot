package commands

import (
	"fmt"
	"time"

	"f1-discord-bot/ergast"

	"github.com/bwmarrin/discordgo"
)

// NextRace performs the actions for the "next" command sent to the bot,
// which informs the user about the next grand prix. The result is a string ready to
// be sent to discord.
func NextRace() (*discordgo.MessageSend, error) {
	// Get next race from the API
	race, err := ergast.RequestNextRace()
	if err != nil {
		return nil, fmt.Errorf("requesting next race to ergast: %v", err)
	}

	// Build message
	var message discordgo.MessageSend
	var gpEmbed discordgo.MessageEmbed
	var practiceEmbed discordgo.MessageEmbed
	var raceEmbed discordgo.MessageEmbed

	message.Embeds = append(message.Embeds, &gpEmbed)
	message.Embeds = append(message.Embeds, &practiceEmbed)
	message.Embeds = append(message.Embeds, &raceEmbed)

	message.Content = "**NEXT RACE INFORMATION**"
	gpEmbed.Title = "Grand Prix Info"

	gpEmbed.Fields = append(gpEmbed.Fields, &discordgo.MessageEmbedField{
		Name:   "Race",
		Value:  race.RaceName,
		Inline: true,
	})

	gpEmbed.Fields = append(gpEmbed.Fields, &discordgo.MessageEmbedField{
		Name:   "Circuit",
		Value:  race.Circuit.CircuitName,
		Inline: true,
	})

	gpEmbed.Fields = append(gpEmbed.Fields, &discordgo.MessageEmbedField{
		Name:   "Location",
		Value:  fmt.Sprintf("%s (%s)", race.Circuit.Location.Locality, race.Circuit.Location.Country),
		Inline: true,
	})

	now := time.Now()

	practiceEmbed.Title = "Practice Sessions"
	if race.FirstPractice != nil {
		practiceEmbed.Fields = append(practiceEmbed.Fields, EmbedForSession("FP1", now, race.FirstPractice, true))
	}

	if race.SecondPractice != nil {
		practiceEmbed.Fields = append(practiceEmbed.Fields, EmbedForSession("FP2", now, race.SecondPractice, true))
	}

	if race.ThirdPractice != nil {
		practiceEmbed.Fields = append(practiceEmbed.Fields, EmbedForSession("FP3", now, race.ThirdPractice, true))
	}

	raceEmbed.Title = "Race Sessions"

	if race.Qualifying != nil {
		raceEmbed.Fields = append(raceEmbed.Fields, EmbedForSession("Qualifying", now, race.Qualifying, true))
	}

	if race.Sprint != nil {
		raceEmbed.Fields = append(raceEmbed.Fields, EmbedForSession("Sprint", now, race.Sprint, true))
	}
	// Parse race time
	raceEmbed.Fields = append(raceEmbed.Fields, EmbedForSession("Race", now, &race.DateTime, true))

	return &message, nil
}

func EmbedForSession(name string, now time.Time, sessionTime *ergast.DateTime, inline bool) *discordgo.MessageEmbedField {
	t := sessionTime.TimeInDefaultLocation()
	delta := t.Sub(now)
	var timeLeftDisplay string
	if delta > 0 {
		timeLeftDisplay = fmt.Sprintf("\n(in %s)", ParseDuration(delta))
	}

	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  fmt.Sprintf("%s%s", t.Format("Mon, 02-Jan-2006\n15:04 MST"), timeLeftDisplay),
		Inline: inline,
	}
}

func EmbedSeparator() *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   "--",
		Value:  "--",
		Inline: false,
	}
}
