package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/mistyyboi/Styrka/config"
	"github.com/mistyyboi/Styrka/events"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	configuration := config.GetConfig()


	dg, err := discordgo.New("Bot " + configuration.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}


	dg.AddHandler(events.HandleMessages)
	dg.AddHandler(events.HandleReady)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}
