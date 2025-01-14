package insurance

import (
	"log"

	"github.com/ozonmp/omp-bot/internal/app/commands/insurance/apartment"
	"github.com/ozonmp/omp-bot/internal/app/path"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type InsuranceCommander struct {
	bot                *tgbotapi.BotAPI
	apartmentCommander Commander
}

func NewInsuranceCommander(
	bot *tgbotapi.BotAPI,
) *InsuranceCommander {

	return &InsuranceCommander{
		bot:                bot,
		apartmentCommander: apartment.NewInsuranceApartmentCommander(bot),
	}
}

func (c *InsuranceCommander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "apartment":
		c.apartmentCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("InsuranceCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *InsuranceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "apartment":
		c.apartmentCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("InsuranceCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
