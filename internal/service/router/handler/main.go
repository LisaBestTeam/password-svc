package handler

import (
	"encoding/json"

	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/sirupsen/logrus"
)

type PasswordHandler struct {
	passwords database.Passwords
	log       *logrus.Logger
}

func NewPasswordHandler(password database.Passwords, logger *logrus.Logger) PasswordHandler {
	return PasswordHandler{passwords: password, log: logger}
}

func (p *PasswordList) Resources() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return marshal
}

type PasswordList struct {
	Data  []database.Password `json:"data"`
	Links Links               `json:"links"`
}

type Links struct {
	Next string `json:"next"`
	Self string `json:"self"`
}

type Error struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Code   int    `json:"code"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}
