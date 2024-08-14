package utils

import "strings"

func ParseCommand(command string) []string {
	var commandParts []string
	parts := strings.Split(command, " ")
	for _, part := range parts {
		commandParts = append(commandParts, strings.ToLower(strings.TrimSpace(part)))
	}
	return commandParts
}

func GetCmdString(commands []string) string {
	var cmd string
	for _, u := range commands {
		cmd = cmd + " " + "'" + u + "'"
	}
	return cmd
}
