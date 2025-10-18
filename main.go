package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"main/config"
	"main/utils"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	cfg := config.Load()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("HI")
	factorialNumber, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Сообщение не является числом",
		})
		return
	}
	factorial := utils.GetFactorial(int64(factorialNumber))
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      fmt.Sprintf("Факториал числа <b>%d</b>\n\n<code>%s</code>", factorialNumber, factorial.String()),
		ParseMode: models.ParseModeHTML,
	})
}
