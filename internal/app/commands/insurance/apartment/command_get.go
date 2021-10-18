package apartment

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *InsuranceApartmentCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("Wrong args", args)
		return
	}

	apartment, err := c.apartmentService.Describe(idx)
	if err != nil {
		log.Printf("Fail to get apartment with idx %d: %v", idx, err)
		return
	}

	var outputMsgText string
	if apartment == nil {
		log.Printf("Apartment #%d not found", idx)
		outputMsgText = fmt.Sprintf("Apartment #%d not found", idx)
	} else {
		outputMsgText = apartment.String()
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.Get: error sending reply message to chat - %v", err)
	}
}
