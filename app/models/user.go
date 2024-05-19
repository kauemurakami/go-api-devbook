package models

import (
	"errors"
	"strings"
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

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.formatter()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome deve ser inserido")
	}
	if user.Email == "" {
		return errors.New("O email deve ser inserido")
	}
	if user.Nick == "" {
		return errors.New("O nick deve ser inserido")
	}
	if user.Pass == "" {
		return errors.New("O password deve ser inserido")
	}

	return nil
}

func (user *User) formatter() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Pass = strings.TrimSpace(user.Pass)
}
