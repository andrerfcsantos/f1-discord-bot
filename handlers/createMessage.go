package handlers

import (
	"fmt"
	"log"
	"strings"

	"f1-discord-bot/commands"

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

	// Process command and figure out the reply to send
	c := ParseCommandArguments(m.Content)

	var message string
	var messageSend *dgo.MessageSend
	var cmdErr error

	switch c.Command {
	case "next":
		messageSend, cmdErr = commands.NextRace()
	case "last":
		message, cmdErr = commands.LastRace()
	case "results":
		message, cmdErr = commands.Results(c.Arguments...)
	case "current":
		message, cmdErr = commands.CurrentSeason()
	case "help":
		message = commands.Help(BOT_PREFIX)
	default:
		cmdErr = fmt.Errorf("command %s not recognized. Type `!f1 help` for a full list of commands available", c.Command)
	}

	if cmdErr != nil {
		message = fmt.Sprintf("Ups, seems like there was a problem executing the command: %v", cmdErr)
	}

	if message != "" {
		messageSend = &dgo.MessageSend{
			Content: message,
		}
	}

	// Send the message
	_, sendErr := s.ChannelMessageSendComplex(m.ChannelID, messageSend)
	if sendErr != nil {
		log.Printf("error sending message to discord: %v", sendErr)
	}

	log.Printf("Guild: %v | Author: %v(%v) | Command: %v | CmdErr: %v | SendErr: %v", m.GuildID, m.Author.ID, m.Author.Username, m.Content, cmdErr, sendErr)
}
