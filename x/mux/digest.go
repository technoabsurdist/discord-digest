package mux

import (
	"os"
	"log"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

/* Out of an array of message structs, create a formatted message and apply logic */ 
func digestCreator() {
	// m.sortByTime()
	// m.sortByContentLength()
	f, err := os.Create("data.md")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString("# Welcome to Your Daily Digest, by the daily digest bot\n")

	// Random gif cause why not 

	// Top news of the day 

	// HackerNews top articles

	// Important stock (faangs, indicators, etc...)  and crypto (btc, eth, sol) prices (including sps, etc.)

	// 
	fmt.Println("done")
}

/* Defines what happens when Digest function is called, which is called when the
bot sees '!digest' in specified channel */ 
func (m *Mux) Digest(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	digestCreator()	
}

