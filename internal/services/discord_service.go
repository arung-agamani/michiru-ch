package services

import (
	"fmt"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Singleton instance and sync control
var (
	instance *DiscordService
	once     sync.Once
)

type DiscordService struct {
	session *discordgo.Session
}

func NewDiscordService() (*DiscordService, error) {
	var err error

	once.Do(func() {
		botToken := os.Getenv("DISCORD_BOT_TOKEN")
		if botToken == "" {
			err = fmt.Errorf("DISCORD_BOT_TOKEN environment variable is not set")
			return
		}
		dg, initErr := discordgo.New("Bot " + botToken)
		if initErr != nil {
			err = fmt.Errorf("failed to create Discord session: %w", initErr)
			return
		}

		initErr = dg.Open()
		if initErr != nil {
			err = fmt.Errorf("failed to open Discord session: %w", initErr)
			return
		}

		instance = &DiscordService{session: dg}
	})

	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (d *DiscordService) SendMessage(channelID, message string) error {
	_, err := d.session.ChannelMessageSend(channelID, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (d *DiscordService) Close() {
	if d.session != nil {
		d.session.Close()
	}
}
