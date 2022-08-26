package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// config settings
type Config struct {
	TOKEN string `mapstructure:"TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

// global err and token
var err error
var token string

func init() {
	// Load token from env file
	config, err := LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	// set our bot's token
	token = config.TOKEN

}
func main() {

	botToken := fmt.Sprintf("Bot %s", token)

	// create the discordgo session
	var Session, _ = discordgo.New(botToken)

	if Session.Token == "" {
		log.Println("You must provide a Discord authenticaton token")
		return
	}

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
