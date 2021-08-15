package core

import "github.com/bwmarrin/discordgo"

type CommandMetadata struct {
	Name string
	Description string
	Usage string
	Aliases []string
}

type CommandContext struct {
	Args []string
	Message *discordgo.Message
	Session *discordgo.Session
}

type Command interface {
	Metadata() CommandMetadata


	Exec(c *CommandContext) error

}

var Commands = make(map[string]Command)
