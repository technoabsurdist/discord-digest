package mux

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

/* Defines what happens when Digest function is called, which is called when the
bot sees '!digest' in specified channel */ 
func (m *Mux) Digest(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	_, err := ds.ChannelMessageSend(dm.ChannelID, "This is your digest")
	if err != nil {
		log.Println(err)
	}
}