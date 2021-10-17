package apartment

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *InsuranceApartmentCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	text := fmt.Sprintf("Get %d. You wrote: : %s", id, inputMsg.Text)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	c.bot.Send(msg)
	//panic("implement me")
}
