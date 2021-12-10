package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/lisabestteam/password-svc/internal/service/router/request"
)

func (p PasswordHandler) GetPasswordReceiver(w http.ResponseWriter, r *http.Request) {
	log := p.log.WithField("handler", "receiver")

	requestReceiver, err := request.GetPasswordReceiver(r)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = requestReceiver.Validate(); err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q := p.passwords.New().Pagination(requestReceiver.Pagination)

	passwords, err := q.SelectByReceiver(requestReceiver.Receiver)
	if err != nil {
		log.WithError(err).Error("failed to get password by receiver")
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonapi.MarshalPayload(w, passwords)
}
