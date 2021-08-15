package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func HandleReady(_ *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("%s is ready\n", r.User.Username)
}
