package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create a new Discord session
	dg, err := discordgo.New("Bot ")
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Register the commandCreate handler
	dg.AddHandler(commandCreate)

	// Open a connection to the Discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	// Wait here until CTRL+C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}

func commandCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.GuildID != "" {
		return
	}

	if m.Content[0] == '/' {
		cmd := m.Content[1:]
		switch cmd {
		case "help":
			s.ChannelMessageSend(m.ChannelID, `I write an "about me" section by inspiring from your LinkedIn profile.

Commands
/help
/start
/linkedin [your linekedin profile URL]`)
		case "linkedin":
			s.ChannelMessageSend(m.ChannelID, "This command can only be used in a private DM.")
		case "start":
			s.ChannelMessageSend(m.ChannelID, "Enter your linkedin profile URL")
		}
	}
}
