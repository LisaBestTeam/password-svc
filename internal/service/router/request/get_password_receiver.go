package request

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/schema"
	"github.com/lisabestteam/password-svc/internal/database"
)

type ReceiverRequest struct {
	Receiver string `schema:"-"`
	database.Pagination
}

func (r ReceiverRequest) Validate() error {
	return validation.Errors{
		"limit":    validation.Validate(r.Limit),
		"page":     validation.Validate(r.Page),
		"receiver": validation.Validate(r.Receiver),
	}.Filter()
}

func GetPasswordReceiver(r *http.Request) (ReceiverRequest, error) {
	var request ReceiverRequest

	if err := schema.NewDecoder().Decode(&request, r.URL.Query()); err != nil {
		return ReceiverRequest{}, err
	}

	request.Receiver = chi.URLParam(r, "receiver")

	return request, request.Validate()
}
