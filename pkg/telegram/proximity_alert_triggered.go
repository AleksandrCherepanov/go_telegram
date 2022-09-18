package telegram

import "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/user"

type ProximityAlertTriggered struct {
	Traveler user.User `json:"traveler"`
	Watcher  user.User `json:"watcher"`
	Distance int64     `json:"distance"`
}
