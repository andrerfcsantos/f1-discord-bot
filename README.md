# Formula 1 Discord Bot

[![Go Report Card](https://goreportcard.com/badge/github.com/andrerfcsantos/f1-discord-bot)](https://goreportcard.com/report/github.com/andrerfcsantos/f1-discord-bot)
[![GoDoc](https://godoc.org/github.com/andrerfcsantos/f1-discord-bot?status.svg)](https://godoc.org/github.com/andrerfcsantos/f1-discord-bot)

Discord bot that posts information about Formula 1 on discord by user request.

## Invite the bot to your Discord server

The easiest way to use bot is to just [invite the bot](https://discordapp.com/api/oauth2/authorize?client_id=595651486923358238&permissions=67632192&scope=bot) to your Discord server and start using the commands.

See below if you want to run the bot on your own machine/server.

## Usage

The users of the server can send commands to the bot using the following:

```
usage: !f1 [command] [command_args...]
Available commands:
    - help - shows this message
    - next - shows information about the next race
    - last - shows information about the last race
    - current - shows races for the current season
    - results - shows information about results
        - results circuit <circuit> - shows historical information about the winners at a given circuit for the last years
        - results driver <driver> - shows last results for a driver
```

The bot will reply in the same channel the command was executed.

## Running the bot on your own server/machine

The easiest way to run the bot is to invite the bot to your Discord server like previously mentioned. By doing that, you are using an instance of the bot running in the cloud.

This section is intended for advanced users who want to run the bot on their own machine or on their own server.

### Prerequisites

* [Go](https://golang.org/dl/) (1.11+ required, since this project uses modules)
* [An Application and Bot User](https://discordapp.com/developers/applications) with permissions to read write messages in the server you want the bot to run.

A binary distribution solution will be considered in the future in order to remove the Go requirement. A docker version is also planned.

### Instalation

* Clone the repository
* `$ cd <repo_root_folder>`
* `$ go build`

    This will create a binary file named `f1-discord-bot` (linux/mac) or `f1-discord-bot.exe` (Windows) that you can use to run the bot. If you want, you can use this binary to deploy the bot on places where Go is not installed. To compile the bot to a different architecture and OS of the machine you're building on, set the `GOOS` and `GOARCH` environment variables accordingly before performing this command. [More info on cross-compiling with Go](https://www.yellowduck.be/posts/cross-compile/)

* Set the `DISCORD_BOT_TOKEN` environemnt variable with your discord bot token. Alternatively, you can also pass this token to the program using the flag `-bot-token`  (see below).

### Run the program

Assuming you already set up `DISCORD_BOT_TOKEN`, run:

* `$ ./f1-discord-bot` (linux/mac)
* `$ f1-discord-bot.exe` (windows)

That's it! The bot should now be running.

If you don't want to use a global environment variable with your token, or you plan to run several instances, you can also define the bot token for each run:

* `$ DISCORD_BOT_TOKEN=<YOUR_BOT_TOKEN> ./f1-discord-bot` (linux/mac)

    or passing it via flag:

* `$ ./f1-discord-bot -bot-token <YOUR_BOT_TOKEN>` (linux/mac)
* `$ f1-discord-bot.exe -bot-token <YOUR_BOT_TOKEN>` (windows)

## Acknowledgements

The information provided by this bot comes from the [Ergast API](https://ergast.com/mrd/).
