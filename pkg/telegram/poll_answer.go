package telegram

import "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/user"

type PollAnswer struct {
	PollId    string    `json:"poll_id"`
	User      user.User `json:"user"`
	OptionIds []int64   `json:"option_ids"`
}
