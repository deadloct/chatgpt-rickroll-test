package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const RickRollCMDName = "rickroll"

func main() {
	// Set up the Discord session
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register a handler for the /rickroll command
	discord.AddHandler(handleSlashCommand)

	// Open a websocket connection to Discord and begin listening
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	cmd := &discordgo.ApplicationCommand{
		Name:        RickRollCMDName,
		Description: "show the rickroll youtube link",
	}
	if _, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", cmd); err != nil {
		panic(err)
	}

	// Keep the program running
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func handleSlashCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.ApplicationCommandData().Name == RickRollCMDName {
		// Send the user the official music video for "Never Gonna Give You Up"
		response := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			},
		}
		session.InteractionRespond(interaction.Interaction, &response)
	}
}
