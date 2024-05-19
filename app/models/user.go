package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omityempty"`
	Name      string    `json:"name,omityempty"`
	Nick      string    `json:"nick,omityempty"`
	Email     string    `json:"email,omityempty"`
	Pass      string    `json:"pass,omityempty"`
	CreatedAt time.Time `json:"created_at,omityempty"`
}
