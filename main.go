package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"technoabsurdist/digest/util"
	"github.com/bwmarrin/discordgo"
)

// global err and token
var err error
var token string
// create the discordgo session
var Session *discordgo.Session


func init() {
	// Load token from env file
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	// set our bot's token
	token = config.TOKEN
	botToken := fmt.Sprintf("Bot %s", token)

	if Session.Token == "" {
		log.Println("You must provide a Discord authenticaton token")
		return
	}
	Session, _ = discordgo.New(botToken)
}

func main() {
	// Open a websocket connection to Discord
	err = Session.Open()
	if err != nil {
		log.Printf("Error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	// wait for Ctrl-C
	log.Printf("Now running. Press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// connect to Discord

	// Log out messages

	// exit
}
