package bot

import (
	"context"
	"log"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CtxKey string

func RunBot(ctx context.Context, cfg *Config) error {
	log.Println("Starting repost bot")
	cfg.Log()
	opts := []tgbot.Option{
		tgbot.WithAllowedUpdates([]string{"channel_post"}),
		tgbot.WithDefaultHandler(defaultHandler),
	}

	bot, err := tgbot.New(cfg.ApiToken, opts...)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, CtxKey("config"), cfg)

	bot.Start(ctx)
	return nil
}

func defaultHandler(ctx context.Context, b *tgbot.Bot, update *models.Update) {
	cfg := ctx.Value(CtxKey("config")).(*Config)
	if update.ChannelPost != nil {
		if update.ChannelPost.Chat.Username != cfg.ChannelName {
			return
		}
		handleChannelPost(ctx, b, update)
	} else {
		log.Println("Not handling update")
	}
}

func handleChannelPost(ctx context.Context, b *tgbot.Bot, update *models.Update) {
	post := update.ChannelPost
	cfg := ctx.Value(CtxKey("config")).(*Config)
	chatID := cfg.ChatID
	threadID := cfg.ThreadID

	fwdParams := &tgbot.ForwardMessageParams{
		ChatID:          chatID,
		MessageThreadID: threadID,
		FromChatID:      post.Chat.ID,
		MessageID:       post.ID,
		ProtectContent:  true,
	}

	_, err := b.ForwardMessage(ctx, fwdParams)
	if err != nil {
		log.Printf("error forwarding message: %v", err)
	}
}
