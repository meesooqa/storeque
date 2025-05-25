package commands

import (
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/google/uuid"
)

func (o *Commander) Tst(inputMessage *tgbotapi.Message) {
	// Ошибка при создании счета для оплаты: Bad Request: STARS_INVOICE_INVALID
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
	fmt.Println("Invoice sent successfully:", sentInvoice)
}
