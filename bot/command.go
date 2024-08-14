package bot

import (
	"fmt"

	"github.com/shahinrahimi/teleprompt/utils"
)

const (
	COMMAND_REGISTER      string = "/start"
	COMMAND_VIEW_ME       string = "/info"
	COMMAND_DELETE_ME     string = "/kick"
	COMMAND_CREATE_PROMPT string = "/add"
	COMMAND_DELETE_PROMPT string = "/delete"

	COMMAND_VIEW_USERS string = "/viewusers"
	COMMAND_KICK_USER  string = "/kickuser"
	COMMAND_BAN_USER   string = "/banuser"
)

func (t *TelegramBot) handleNewUserCommand(userID int64, command string) {
	commandParts := utils.ParseCommand(command)
	switch {
	case commandParts[0] == COMMAND_REGISTER:
		break
	default:
		t.handleSendUsageForNewUser(userID)
	}
}

func (t *TelegramBot) handleUserCommand(userID int64, command string) {
	commandParts := utils.ParseCommand(command)
	switch {
	case commandParts[0] == COMMAND_VIEW_ME:
		break
	case commandParts[0] == COMMAND_DELETE_ME:
		break
	case commandParts[0] == COMMAND_CREATE_PROMPT:
		break
	case commandParts[0] == COMMAND_DELETE_PROMPT:
		break
	default:
		t.handleSendUsageForUser(userID)
	}
}

func (t *TelegramBot) handleAdminCommand(userID int64, command string) {
	commandParts := utils.ParseCommand(command)
	switch {
	case commandParts[0] == COMMAND_VIEW_ME:
		break
	case commandParts[0] == COMMAND_CREATE_PROMPT:
		break
	case commandParts[0] == COMMAND_DELETE_PROMPT:
		break
	case commandParts[0] == COMMAND_VIEW_USERS:
		break
	case commandParts[0] == COMMAND_KICK_USER:
		break
	case commandParts[0] == COMMAND_BAN_USER:
		break
	default:
		t.handleSendUsageForAdmin(userID)
	}
}

func (t *TelegramBot) handleSendUsageForNewUser(userId int64) {
	usageCommands := []string{
		COMMAND_REGISTER,
	}
	cmdString := utils.GetCmdString(usageCommands)
	msgStr := fmt.Sprintf("Usage\n commands: %s", cmdString)
	t.sendMessage(userId, msgStr)
}

func (t *TelegramBot) handleSendUsageForUser(userId int64) {
	usageCommands := []string{
		COMMAND_VIEW_ME,
		COMMAND_DELETE_ME,
		COMMAND_CREATE_PROMPT,
		COMMAND_DELETE_PROMPT,
	}
	cmdString := utils.GetCmdString(usageCommands)
	msgStr := fmt.Sprintf("Usage\n commands: %s", cmdString)
	t.sendMessage(userId, msgStr)
}

func (t *TelegramBot) handleSendUsageForAdmin(userId int64) {
	usageCommands := []string{
		COMMAND_VIEW_ME,
		COMMAND_CREATE_PROMPT,
		COMMAND_DELETE_PROMPT,
		COMMAND_VIEW_USERS,
		COMMAND_KICK_USER,
		COMMAND_BAN_USER,
	}
	cmdString := utils.GetCmdString(usageCommands)
	msgStr := fmt.Sprintf("Usage\n commands: %s", cmdString)
	t.sendMessage(userId, msgStr)
}
