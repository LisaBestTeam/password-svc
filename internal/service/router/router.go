package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lisabestteam/password-svc/internal/service/router/handler"
)

func (r router) NewRouter() chi.Router {
	router := chi.NewRouter()

	newHandler := handler.NewHandler()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		)

	router.Route("/integrations/password", func(r chi.Router) {
		r.Get("/receiver/{address}", newHandler.GetPasswordReceiver)
		r.Get("/sender/{address}", newHandler.GetPasswordSender)
	})

	return router
}
