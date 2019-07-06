package main

import (
	"f1-discord-bot/handlers"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	dgo "github.com/bwmarrin/discordgo"
)

// BOT_TOKEN represents the discord authentication token
var BOT_TOKEN string

var session *dgo.Session

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {

	// Discord Authentication Token
	BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	if BOT_TOKEN == "" {
		flag.StringVar(&BOT_TOKEN, "bot-token", "", "Discord Authentication Token")
	}
	flag.Parse()
}

func main() {
	var err error

	if BOT_TOKEN == "" {
		log.Print("No bot token specified. Please specify one using the DISCORD_BOT_TOKEN environment variable or the -bot-token flag.")
		return
	}

	session, err = dgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		log.Printf("error getting new session: %v", err)
		return
	}

	err = session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		return
	}
	defer session.Close()

	session.UpdateStatus(0, "!f1 help")
	session.AddHandler(handlers.CreateMessage)

	// Wait for a CTRL-C
	log.Printf("It's lights out and away we go! Bot now running. (CTRL-C to exit)")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
