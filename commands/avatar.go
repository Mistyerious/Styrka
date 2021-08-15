package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/mistyyboi/Styrka/core"
)


type AvatarCommand struct{}

func (h *AvatarCommand) Metadata() core.CommandMetadata {

	return core.CommandMetadata{
		Name: "avatar",
		Description: "Gives you the avatar of the mentioned user",
		Usage: "avatar [@user]",
		Aliases: []string{"avatar", "a"},
	}

}

func (h *AvatarCommand) Exec(c *core.CommandContext) error {

	var user *discordgo.User
	if len(c.Message.Mentions) == 0 {
		user = c.Message.Author
	} else { user = c.Message.Mentions[0] }


	_, err := c.Session.ChannelMessageSend(c.Message.ChannelID, user.Username); if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func init() {
	core.Commands["avatar"] = &AvatarCommand{}
	core.Commands["a"] = &AvatarCommand{}
}