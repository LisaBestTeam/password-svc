package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/lisabestteam/password-svc/internal/service/router/request"
)

func (p PasswordHandler) GetPasswordSender(w http.ResponseWriter, r *http.Request) {
	log := p.log.WithField("handler", "receiver")

	requestSender, err := request.GetPasswordSender(r)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = requestSender.Validate(); err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q := p.passwords.New().Pagination(requestSender.Pagination)

	passwords, err := q.SelectBySender(requestSender.Sender)
	if err != nil {
		log.WithError(err).Error("failed to get password by sender")
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	jsonapi.MarshalPayload(w, passwords)
}
