package mux

import (
	"os"
	"log"
	"fmt"
	"time"
	"sort"
	"github.com/bwmarrin/discordgo"
)

type MessageContainer struct {
	author string
	content string
	timestamp time.Time
	channelID string
}

type MessagesArr []MessageContainer

/* Sort messages by calculated time for user */ 
func (m MessagesArr) sortByTime() {
	// sort by time
	sort.Slice(m, func(i, j int) bool {
		return m[i].timestamp.Unix() > m[j].timestamp.Unix()
	})	
}

func (m MessagesArr) sortByContentLength() {
	// sort by length of content
	sort.Slice(m, func(i, j int) bool {
		m_i := len([]rune(m[i].content)) 
		m_j := len([]rune(m[j].content))
		return m_i > m_j
	})	
}

/* Out of an array of message structs, create a formatted message and apply logic */ 
func digestCreator(m MessagesArr) string {
	m.sortByTime()
	m.sortByContentLength()
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	totalString := ""
	for _, message := range m {
		stringToAdd := fmt.Sprintf("%s: %s: %s: ", message.timestamp, message.author, message.content)
		_, err2  := f.WriteString(stringToAdd)
		totalString += stringToAdd
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Println("done")
	return totalString
}

/* Defines what happens when Digest function is called, which is called when the
bot sees '!digest' in specified channel */ 
func (m *Mux) Digest(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	ds.ChannelMessageSend(dm.ChannelID, "Putting your digest together...")
	_, serverIDs := channelNameAndIDs(ds)
	messages := getMessages(ds, serverIDs)
	returnMessage := digestCreator(messages)	
	ds.ChannelMessageSend(dm.ChannelID, returnMessage) // checking that reaching endpoint
}

/* get channel names and ids from a specific guild (server) */ 
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

	return nameServers, serverIDs
}

/* Function to get past 100 messages from array of channel IDs */   
func getMessages(s *discordgo.Session, ids []string) ([]MessageContainer) {
	// print all messages from all channels
	messagesContainer := MessagesArr{}
	for _, id := range ids {
		messages, _ := s.ChannelMessages(id, 100, "", "", "")
		for _, message := range messages {
			m := MessageContainer{message.Author.Username, message.Content, message.Timestamp, message.ChannelID}
			messagesContainer = append(messagesContainer, m)		
		}
	}

	return messagesContainer
}
