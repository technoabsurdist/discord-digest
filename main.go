package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"technoabsurdist/digest/util"
	"technoabsurdist/digest/x/mux"
	"github.com/bwmarrin/discordgo"
)

// global err and token
var err error
var token string
var Session *discordgo.Session
var Router *mux.Mux

func init() {
	// Load token from env file
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// set our bot's token
	token = config.TOKEN
	botToken := fmt.Sprintf("Bot %s", token)

	Session, _ = discordgo.New(botToken)
	if Session.Token == "" {
		log.Println("You must provide a Discord authenticaton token")
		return
	}

	Router = mux.New()
	Session.AddHandler(Router.OnMessageCreate)
	
	// Routes 
	Router.Route("!help", "Help menu", Router.Help)
	Router.Route("!digest", "Digest", Router.Digest)
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
}
