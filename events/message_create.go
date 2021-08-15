package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mistyyboi/Styrka/commands"
	"github.com/mistyyboi/Styrka/config"
	"github.com/mistyyboi/Styrka/core"
	"strings"
)


func HandleMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	configuration := config.GetConfig()

	if !strings.HasPrefix(m.Content, configuration.Prefix) {
		return
	}

	args := strings.Split(strings.ToLower(m.Content), " ")

	cmd, exists := core.Commands[strings.TrimPrefix(args[0], configuration.Prefix)]

	if !exists { return }

	context := new(core.CommandContext)
	context.Args = args
	context.Message = m.Message
	context.Session = s


	if err := cmd.Exec(context); err != nil {
		fmt.Println(err)
		return
	}

}
