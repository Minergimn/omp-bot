package apartment

import (
	"encoding/json"
	"fmt"
	"log"

	serviceApartment "github.com/ozonmp/omp-bot/internal/service/insurance/apartment"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *InsuranceApartmentCommander) New(inputMessage *tgbotapi.Message) {
	parsedApartment := serviceApartment.Apartment{}
	args := inputMessage.CommandArguments()
	err := json.Unmarshal([]byte(args), &parsedApartment)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.New: "+
			"error reading json data for type Apartment from "+
			"input string %v - %v", args, err)
		return
	}

	var outputMsgText string
	id, err := c.apartmentService.Create(parsedApartment)
	if err != nil {
		outputMsgText = fmt.Sprintf("Fail to create apartment %s: %v", parsedApartment.String(), err)
		log.Print(outputMsgText)
	} else {
		outputMsgText = fmt.Sprintf("Apartment was added with id %d", id)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.New: error sending reply message to chat - %v", err)
	}
}
