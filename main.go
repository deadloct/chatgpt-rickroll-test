package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

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

	// Keep the program running
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	http.ListenAndServe(":8080", nil)
}

func handleSlashCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Data.Name == "rickroll" {
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
