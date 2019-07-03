package commands

import "fmt"

// Help performs the actions for the "help" command sent to the bot,
// which informs the user about the usage and commands available
func Help(prefix string) string {
	return fmt.Sprintf("usage: %s [command] [command_args...]\n", prefix) +
		`Available commands:
    help - shows this message
	next - shows information about the next race
	last - shows information about the last race
`
}
