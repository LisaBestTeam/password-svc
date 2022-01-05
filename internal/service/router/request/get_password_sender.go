package request

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/schema"
	"github.com/lisabestteam/password-svc/internal/database"
)

type SenderRequest struct {
	Sender string `schema:"-"`
	database.Pagination
}

func (r SenderRequest) Validate() error {
	return validation.Errors{
		"limit":  validation.Validate(r.Limit),
		"page":   validation.Validate(r.Page),
		"sender": validation.Validate(r.Sender),
	}.Filter()
}

func GetPasswordSender(r *http.Request) (SenderRequest, error) {
	var request SenderRequest

	if err := schema.NewDecoder().Decode(&request, r.URL.Query()); err != nil {
		return SenderRequest{}, err
	}

	if request.Limit == 0 {
		request.Limit = 15
	}
	request.Sender = chi.URLParam(r, "address")

	return request, request.Validate()
}
