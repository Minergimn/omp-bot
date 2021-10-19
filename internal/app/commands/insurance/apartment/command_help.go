package apartment

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__insurance__apartment - help\n"+
			"/get__insurance__apartment {id} - get by apartment id\n"+
			"/list__insurance__apartment - list products\n"+
			"/list__insurance__apartment {limit} - list products with pagination\n"+
			"/delete__insurance__apartment {id} - delete by apartment id\n"+
			"/new__insurance__apartment {json} - add a new apartment\n"+
			"/edit__insurance__apartment {json} - edit apartment (implemented)\n",
	)

	c.bot.Send(msg)
}
