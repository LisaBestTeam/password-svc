package handler

import (
	"encoding/json"
	"github.com/lisabestteam/password-svc/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (p PasswordHandler) GetPasswordReceiver(w http.ResponseWriter, r *http.Request) {
	receiver := chi.URLParam(r, "receiver")

	q := p.passwords.New()
	log := p.log

	passwords, err := q.SelectByReceiver(receiver)
	if err != nil {
		log.WithError(err).Error("failed to get password by receiver")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: []Error{{Code: http.StatusInternalServerError, Title: "Database return error", Detail: err.Error()}}})
		return
	}

	if passwords == nil {
		passwords = make([]database.Password, 0)
	}

	json.NewEncoder(w).Encode(PasswordList{Data: passwords})

}
