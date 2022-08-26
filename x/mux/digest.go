package mux

import (
	// "fmt"
	"log"
	// "net/http"
	"github.com/bwmarrin/discordgo"
)

/* Defines what happens when Digest function is called, which is called when the
bot sees '!digest' in specified channel */ 
func (m *Mux) Digest(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	ds.ChannelMessageSend(dm.ChannelID, "Reached your digest") // checking that reaching endpoint

	// myUrl  := "testing"
	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", myUrl, nil)
	// res, _ := client.Do(req)
	// bod := res.Body
	// fmt.Println("%v\n", bod)
	nameServers, serverIDs := channelNameAndIDs(ds)
	log.Println("name servers: %v\n server IDs: %v\n", nameServers, serverIDs)
}

func channelNameAndIDs(s *discordgo.Session) ([]string, []string) {
    // Loop through each guild in the session
	nameServers := []string{}
	serverIDs := []string{}

    for _, guild := range s.State.Guilds {

        // Get channels for this guild
        channels, _ := s.GuildChannels(guild.ID)

        for _, c := range channels {
            // Check if channel is a guild text channel and not a voice or DM channel
            if c.Type != discordgo.ChannelTypeGuildText {
                continue
            }
			// add server names and ids to two arrays respectively 
			nameServers = append(nameServers, c.Name)
			serverIDs = append(serverIDs, c.ID)
        }
    }

	// print all messages from all servers
	for _, id := range serverIDs {
		messages, _ := s.ChannelMessages(id, 100, "", "", "")
		for _, message := range messages {
			log.Println("Messages: %s\n", message.Content)
		}
	}
	return nameServers, serverIDs
}