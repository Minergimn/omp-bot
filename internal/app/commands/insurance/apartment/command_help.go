package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__insurance__apartment - help\n"+
			"/get__insurance__apartment {id} - get by apartment id\n"+
			"/list__insurance__apartment - list products\n"+
			"/delete__insurance__apartment - delete by apartment id\n"+
			"\nnot implemented yet:\n"+
			"/new__insurance__apartment - add new insurance for apartment\n"+
			"/edit__insurance__apartment - edit insurance for apartment\n",
	)

	c.bot.Send(msg)
}
