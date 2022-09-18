package telegram

import "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/user"

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            user.User `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageId *string   `json:"inline_message_id"`
	Query           string    `json:"query"`
}
