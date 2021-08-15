package commands

import (
	"fmt"
	"github.com/mistyyboi/Styrka/core"
)


type HelpCommand struct{}

func (h *HelpCommand) Metadata() core.CommandMetadata {

	return core.CommandMetadata{
		Name: "help",
		Description: "Gives you the help menu for the bot",
		Usage: "help [command]",
		Aliases: []string{"help", "h"},
	}

}

func (h *HelpCommand) Exec(c *core.CommandContext) error {

	for _, command := range core.Commands {
		fmt.Println(command.Metadata().Name)
	}

	return nil
}

func init() {
	core.Commands["help"] = &HelpCommand{}
	core.Commands["h"] = &HelpCommand{}
}