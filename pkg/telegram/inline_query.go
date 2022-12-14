package telegram

import "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/user"

type InlineQuery struct {
	Id       string    `json:"id"`
	From     user.User `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType *string   `json:"chat_type"`
	Location *Location `json:"location"`
}
