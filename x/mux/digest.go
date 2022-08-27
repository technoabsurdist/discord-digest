package mux

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/bwmarrin/discordgo"
	"github.com/caser/gophernews"
)

/* Out of an array of message structs, create a formatted message and apply logic */ 
func digestCreator(dm *discordgo.Message) {
	// m.sortByTime()
	// m.sortByContentLength()
	f, err := os.Create("data.md")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("# Welcome to Your Daily Digest, @%s! \n", dm.Author.Username))
	year, month, day := time.Now().Date()
	f.WriteString(fmt.Sprintf("%s, %d, %d\n", month, day, year))
	f.WriteString(fmt.Sprint("![robot news](https://media.giphy.com/media/LP0ZqMGRx0azpBhBgN/giphy.gif)\n"))
	client := gophernews.NewClient()
	res, err := client.GetTop100()
	resTop10 := res[:10]

	_, err = f.WriteString("\n\n## → 📰 Top HackerNews Today</u>\n")
	for _, topStory := range resTop10 {
		story, err := client.GetStory(topStory)
		if err != nil {
			continue
		}
		titleFormatted := fmt.Sprintf("\n**%s**\n\t", story.Title)
		urlFormatted := fmt.Sprintf("<u>%s</u>\n", story.URL)
		write := titleFormatted + urlFormatted

		_, err = f.WriteString(write)
		if err != nil {
			continue
		}
	}

	// Important stock (faangs, indicators, etc...)  and crypto (btc, eth, sol) prices (including sps, etc.)
	_, err = f.WriteString("\n\n## → 📈 Top Stock Indicators Today \n")

	_, err = f.WriteString("\n\n## → ₿ Top Crypto Indicators Today\n")
}

/* Defines what happens when Digest function is called, which is called when the
bot sees '!digest' in specified channel */ 
func (m *Mux) Digest(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	digestCreator(dm)	
}

