package telegram

import "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/user"

type CallbackQuery struct {
	Id              string    `json:"id"`
	From            user.User `json:"from"`
	Message         *Message  `json:"message"`
	InlineMessageId *string   `json:"inline_message_id"`
	ChatInstance    string    `json:"chat_instance"`
	Data            *string   `json:"data"`
	GameShortName   *string   `json:"game_short_name"`
}
