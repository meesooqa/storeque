package commands

import (
	"context"
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/google/uuid"

	"github.com/meesooqa/storeque/common/app"
	"github.com/meesooqa/storeque/common/lang"
)

type BuyHandler struct {
	BaseHandler
}

func NewBuyHandler(appDeps app.App, bot *tgbotapi.BotAPI) *BuyHandler {
	return &BuyHandler{
		BaseHandler{
			bot:     bot,
			appDeps: appDeps,
		},
	}
}

func (o BuyHandler) GetName() string {
	return "buy"
}

func (o BuyHandler) GetDescription(loc lang.Localization) string {
	return loc.Localize(fmt.Sprintf("tg.cmd.%s.description", o.GetName()), nil)
}

func (o *BuyHandler) Handle(ctx context.Context, loc lang.Localization, inputMessage *tgbotapi.Message) {
	// TODO o.appDeps.Lang.Localize()
	// TODO promocode
	product := struct {
		ID          string
		Name        string
		Description string
	}{
		ID:          uuid.NewString(),
		Name:        "Product Name",
		Description: "Product Description",
	}
	startParameter := uuid.NewString()
	// https://core.telegram.org/bots/api#sendinvoice
	prices := []tgbotapi.LabeledPrice{{
		//Label:  "Full Price",
		Label:  "XTR",
		Amount: 1,
	}}
	// suggestedTipAmounts := []int{10, 100, 500, 1000}
	suggestedTipAmounts := []int{}
	invoice := tgbotapi.NewInvoice(
		inputMessage.Chat.ID,
		product.Name,
		product.Description,
		fmt.Sprintf("product:%d", product.ID),
		"",             // Pass an empty string for payments in Telegram Stars
		startParameter, // fmt.Sprintf("%d", inputMessage.From.ID),
		"XTR",
		prices,
		suggestedTipAmounts,
	)
	sentInvoice, err := o.bot.Send(invoice)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Ошибка при создании счета для оплаты: %v", err))
		o.bot.Send(msg)
		return
	}
	fmt.Printf("Invoice sent successfully: %+v", sentInvoice)
}
