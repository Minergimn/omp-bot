package apartment

import (
	"encoding/json"
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/insurance/apartment"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func getPage(chatId int64, apartments []apartment.Apartment, cursor uint64, offset uint64, bot *tgbotapi.BotAPI) {
	outputMsgText := "Here are all the insured apartments: \n\n"
	var returnedCount int

	for i, a := range apartments {
		outputMsgText += a.String()
		outputMsgText += "\n"
		returnedCount = i + 1
	}

	msg := tgbotapi.NewMessage(chatId, outputMsgText)

	elementsOver := offset-uint64(returnedCount) > 0
	if offset == 0 || elementsOver {
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("InsuranceApartmentCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: offset,
		Cursor: cursor + offset,
	})

	callbackPath := path.CallbackPath{
		Domain:       "insurance",
		Subdomain:    "apartment",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.List: error sending reply message to chat - %v", err)
	}
}
