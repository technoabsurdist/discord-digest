package mux

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/bwmarrin/discordgo"
	"github.com/caser/gophernews"
	"github.com/svarlamov/goyhfin"
)


/* Out of an array of message structs, create a formatted message and apply logic */ 
func digestCreator(ds *discordgo.Session, dm *discordgo.Message) {
	f, err := os.Create("data.md")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("# Welcome to Your Daily Digest, @%s! \n", dm.Author.Username))
	year, month, day := time.Now().Date()
	f.WriteString(fmt.Sprintf("%s, %d, %d\n", month, day, year))
	f.WriteString(fmt.Sprint("![robot news](https://media.giphy.com/media/LP0ZqMGRx0azpBhBgN/giphy.gif)\n"))
	
	getHackerNews(f)
	writeHr(f)
	_, err = f.WriteString("\n\n## → 📈 Top Stock Indicators Today \n")
	getStockIndicators(f)

	writeHr(f)
	_, err = f.WriteString("\n\n## → ₿ Top Crypto Indicators Today\n")
	getCryptoIndicators(f)
	
	// send file as message
	var data *discordgo.MessageSend
	dataFile, err := os.Open("data.md")
	if err != nil {
		log.Fatal(err)
	}


	discordF := &discordgo.File{
		Name:   "data.md",
		ContentType: "text/markdown",
		Reader: dataFile,
	}
	data = &discordgo.MessageSend{
		Content: "",
		Files: []*discordgo.File{discordF},
	}
	ds.ChannelMessageSendComplex(dm.ChannelID, data)
	log.Println("Sent digest")
	f.Close()
}

/* Defines what happens when Digest function is called, which is called when the
bot sees '!digest' in specified channel */ 
func (m *Mux) Digest(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	digestCreator(ds, dm)	
}


func getHackerNews(f *os.File) {
	client := gophernews.NewClient()
	res, _ := client.GetTop100()
	resTop10 := res[:15]

	writeHr(f)
	f.WriteString("\n\n## → 📰 Top HackerNews Today</u>\n")
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
}

func getStockIndicators(f *os.File) {
	// SP500, Dow Jones, Apple, Google, Meta, Twitter, Amazon, Tesla, 	
	symbols := []string {"^GSPC", "^DJI", "AAPL", "META", "TWTR", "AMZN", "TSLA"} 
	for _, symbol := range symbols {
		resp, err := goyhfin.GetTickerData(symbol, goyhfin.OneMonth, goyhfin.OneDay, false)
		if err != nil {
			continue
			fmt.Println("Error fetching Yahoo Finance data:", err)
		}

		f.WriteString(fmt.Sprintf("\n\n### %s\n\n", symbol))
		for ind := range resp.Quotes {
				quote := resp.Quotes[ind]
				f.WriteString(fmt.Sprintf("\t**%s %d:** %f\n", 
					quote.OpensAt.Month(), quote.OpensAt.Day(), quote.High))
		}
	}
}

func getCryptoIndicators(f *os.File) {
	symbols := []string {"BTC-USD", "SOL-USD", "ETH-USD"} 
	for _, symbol := range symbols {
		resp, err := goyhfin.GetTickerData(symbol, goyhfin.OneMonth, goyhfin.OneDay, false)
		if err != nil {
			continue
			fmt.Println("Error fetching Yahoo Finance data:", err)
		}

		f.WriteString(fmt.Sprintf("\n\n### %s\n\n", symbol))
		for ind := range resp.Quotes {
				quote := resp.Quotes[ind]
				f.WriteString(fmt.Sprintf("\t**%s %d:** %f\n", 
					quote.OpensAt.Month(), quote.OpensAt.Day(), quote.High))
		}
	}
}

func writeHr(f *os.File) {
	f.WriteString("\n\n--------------------------------------\n")
}