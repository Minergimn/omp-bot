package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) List(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "List. You wrote: "+inputMsg.Text)

	c.bot.Send(msg)
	//panic("implement me")
}