package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/insurance/apartment"
	"log"
)

type ApartmentCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)    // return error not implemented
	Edit(inputMsg *tgbotapi.Message)   // return error not implemented
}

type InsuranceApartmentCommander struct {
	bot              *tgbotapi.BotAPI
	apartmentService *apartment.InsuranceApartmentService
}

func NewInsuranceApartmentCommander(bot *tgbotapi.BotAPI) *InsuranceApartmentCommander {

	service := apartment.NewInsuranceApartmentService()

	return &InsuranceApartmentCommander{
		bot:              bot,
		apartmentService: service,
	}
}

func (c *InsuranceApartmentCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("InsuranceApartmentCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *InsuranceApartmentCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}