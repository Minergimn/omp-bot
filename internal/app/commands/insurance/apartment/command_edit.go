package apartment

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/insurance/apartment"
	"log"
)

func (c *InsuranceApartmentCommander) Edit(inputMessage *tgbotapi.Message) {
	parsedApartment := service.Apartment{}
	args := inputMessage.CommandArguments()
	err := json.Unmarshal([]byte(args), &parsedApartment)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.Edit: "+
			"error reading json data for type Apartment from "+
			"input string %v - %v", args, err)
		return
	}

	var outputMsgText string
	err = c.apartmentService.Update(parsedApartment.ApartmentId, parsedApartment)
	if err != nil {
		outputMsgText = fmt.Sprintf("Fail to edit apartment %s: %v", parsedApartment.String(), err)
		log.Print(outputMsgText)
	} else {
		outputMsgText = fmt.Sprintf("Apartment  with id %d was edited", parsedApartment.ApartmentId)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.Edit: error sending reply message to chat - %v", err)
	}
}
