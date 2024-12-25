package bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	ApiToken    string
	ChannelName string
	ChatID      int
	ThreadID    int
}

func (c *Config) Log() {
	log.Printf("BotApi Token: %s", c.ApiToken[:4]+"..."+c.ApiToken[len(c.ApiToken)-3:])
	log.Printf("Channel Name: %s", c.ChannelName)
	log.Printf("Chat ID: %d", c.ChatID)
	log.Printf("Thread ID: %d", c.ThreadID)
}

const (
	tokenEnv       = "REP_BOT_API_TOKEN"
	channelNameEnv = "REP_BOT_CHANNEL_NAME"
	chatIDEnv      = "REP_BOT_CHAT_ID"
	threadIDEnv    = "REP_BOT_THREAD_ID"
)

func ensureEnv(key string) (string, error) {
	value, found := os.LookupEnv(key)
	if !found {
		return "", fmt.Errorf("missing %s environment variable", key)
	}
	return value, nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	token, err := ensureEnv(tokenEnv)
	if err != nil {
		return nil, err
	}
	config.ApiToken = token

	channelName, err := ensureEnv(channelNameEnv)
	if err != nil {
		return nil, err
	}
	config.ChannelName = channelName

	chatIDRaw, err := ensureEnv(chatIDEnv)
	if err != nil {
		return nil, err
	}

	chatId, err := strconv.Atoi(chatIDRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid chat id: %s", chatIDRaw)
	}
	config.ChatID = chatId

	config.ThreadID = -1
	// Can be empty
	threadId := os.Getenv(threadIDEnv)
	if threadId == "" {
		return config, nil
	}
	threadID, err := strconv.Atoi(threadId)
	if err != nil {
		return nil, fmt.Errorf("invalid thread id: %s", threadId)
	}
	config.ThreadID = threadID

	return config, nil
}
