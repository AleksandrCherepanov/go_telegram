package telegram

import "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/user"

type VideoChatParticipantsInvited struct {
	Users []user.User `json:"users"`
}
