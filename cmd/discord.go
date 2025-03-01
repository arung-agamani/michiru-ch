package cmd

import (
	"fmt"
	"michiru/internal/services"

	"github.com/spf13/cobra"
)

var (
	botToken  string
	channelID string
	message   string
)

var discordCmd = &cobra.Command{
	Use:   "discord",
	Short: "Send a message to a Discord channel",
	Run: func(cmd *cobra.Command, args []string) {
		discordService, err := services.NewDiscordService()
		if err != nil {
			fmt.Println("Error initializing Discord service:", err)
			return
		}
		defer discordService.Close()

		err = discordService.SendMessage(channelID, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Println("Message sent successfully!")
	},
}

func init() {
	rootCmd.AddCommand(discordCmd)

	discordCmd.Flags().StringVarP(&channelID, "channel", "c", "", "Discord channel ID (required)")
	discordCmd.Flags().StringVarP(&message, "message", "m", "", "Message to send (required)")

	discordCmd.MarkFlagRequired("channel")
	discordCmd.MarkFlagRequired("message")
}
