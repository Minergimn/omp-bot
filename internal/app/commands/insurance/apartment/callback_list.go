package apartment

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
	Cursor uint64 `json:"cursor"`
}

func (c *InsuranceApartmentCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("InsuranceApartmentCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	apartments, err := c.apartmentService.List(parsedData.Cursor, parsedData.Offset)
	if err != nil {
		log.Printf("Fail to get apartment list")
		return
	}

	getPage(callback.Message.Chat.ID, apartments, parsedData.Cursor, parsedData.Offset, c.bot)
	return

	//for _, a := range apartments {
	//	outputMsgText += a.String()
	//	outputMsgText += "\n"
	//}
	//
	//msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	//
	//if parsedData.Offset == 0 {
	//	_, err = c.bot.Send(msg)
	//	if err != nil {
	//		log.Printf("InsuranceApartmentCommander.List: error sending reply message to chat - %v", err)
	//	}
	//	return
	//}
	//
	//serializedData, _ := json.Marshal(CallbackListData{
	//	Offset: parsedData.Offset,
	//	Cursor: parsedData.Cursor + parsedData.Offset,
	//})
	//
	//callbackPath = path.CallbackPath{
	//	Domain:       "insurance",
	//	Subdomain:    "apartment",
	//	CallbackName: "list",
	//	CallbackData: string(serializedData),
	//}
	//
	//msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	//	tgbotapi.NewInlineKeyboardRow(
	//		tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
	//	),
	//)
	//
	//_, err = c.bot.Send(msg)
	//if err != nil {
	//	log.Printf("InsuranceApartmentCommander.List: error sending reply message to chat - %v", err)
	//}
}
