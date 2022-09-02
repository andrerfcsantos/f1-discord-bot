package handlers

import "strings"

// CommandArguments is a tuple with a command and its arguments
type CommandArguments struct {
	Command   string
	Arguments []string
}

// ParseCommandArguments takes a discord command and splits into its action and arguments.
// The command is assumed to not include any bot prefix.
func ParseCommandArguments(command string) CommandArguments {
	var cleanArgs []string

	parts := strings.Split(strings.TrimSpace(command), " ")

	// Spliting for spaces is not enough to get the clean arguments, since
	// two or more consecutive spaces will create an empty argument.
	// This cycle gets rid of "empty" arguments.
	for _, part := range parts {
		if part != "" {
			cleanArgs = append(cleanArgs, part)
		}
	}

	if len(cleanArgs) == 0 {
		return CommandArguments{}
	}

	return CommandArguments{
		Command:   cleanArgs[0],
		Arguments: cleanArgs[1:],
	}
}
