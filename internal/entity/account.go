package entity

import (
	"fmt"
	"time"
)

type Account struct {
	ID string `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	DisplayName string `json:"display_name"`
	JoinedAt time.Time `json:"joined_at"`
}

func (a *Account) EntityID() ID {
	return ID(fmt.Sprintf("%s:%s", a.ID, a.Username))
} 