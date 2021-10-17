package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) Edit(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Edit. You wrote: "+inputMsg.Text)

	c.bot.Send(msg)
	//panic("implement me")
}
