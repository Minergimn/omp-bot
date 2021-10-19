package apartment

import (
	"encoding/json"
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"

	"github.com/go-telegram-bot-api/telegram-bot-api"
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
}
