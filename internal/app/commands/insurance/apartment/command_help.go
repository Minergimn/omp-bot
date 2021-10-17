package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__insurance__apartment - help\n"+
			"/list__insurance__apartment - list products\n"+
			"/get__insurance__apartment {id} - get by apartment id\n"+
			"/edit__insurance__apartment - edit insurance for apartment\n"+
			"/new__insurance__apartment - add new insurance for apartment\n"+
			"/delete__insurance__apartment - delete by apartment id\n",
	)

	c.bot.Send(msg)
}
