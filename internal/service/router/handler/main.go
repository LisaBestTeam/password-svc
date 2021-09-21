package handler

import (
	"net/http"

	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	GetPasswordReceiver(w http.ResponseWriter, r *http.Request)
	GetPasswordSender(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	passwords database.Passwords
	log       *logrus.Logger
}

func NewHandler() Handler {
	return &handler{}
}
