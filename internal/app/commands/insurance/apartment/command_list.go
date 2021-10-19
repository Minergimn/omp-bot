package apartment

import (
	"log"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) List(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	limit, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("Wrong args, its' ok, just set limit as 0", args)
	}

	apartments, err := c.apartmentService.List(0, limit)
	if err != nil {
		log.Printf("Fail to get apartment list")
		return
	}

	getPage(inputMessage.Chat.ID, apartments, 0, limit, c.bot)
}
