package handlers

import (
	"f1-discord-bot/commands"
	"fmt"
	"log"
	"strings"

	dgo "github.com/bwmarrin/discordgo"
)

// BOT_PREFIX is the prefix for any command sent to the bot
const BOT_PREFIX string = "!f1"

// CreateMessage handles a message coming from discord
func CreateMessage(s *dgo.Session, m *dgo.MessageCreate) {

	m.Content = strings.TrimSpace(m.Content)
	// Check if the message is intended for this bot
	if !strings.HasPrefix(m.Content, BOT_PREFIX) {
		return
	}

	if m.Content == BOT_PREFIX {
		// User called the bot but didn't specify a command,
		// assume help command
		m.Content = BOT_PREFIX + " help"
	}

	// Skip the prefix
	m.Content = m.Content[len(BOT_PREFIX)+1:]
	log.Printf("Processing command: %v", m.Content)

	// Process command and figure out the reply to send
	c := ParseCommandArguments(m.Content)

	var message string
	var err error

	switch c.Command {
	case "next":
		message, err = commands.NextRace()
		if err != nil {
			message = fmt.Sprintf("Ups, seems like there was a problem executing the command. The error reported is: %v", err)
		}
	case "help":
		message = commands.Help(BOT_PREFIX)
	}

	// Send the message
	_, err = s.ChannelMessageSend(m.ChannelID, message)

	if err != nil {
		log.Printf("error sending message to discord: %v", err)
		return
	}

}
