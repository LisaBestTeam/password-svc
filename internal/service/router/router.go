package router

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/lisabestteam/password-svc/internal/service/router/handler"
)

func (r server) NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
	)

	password := handler.NewPasswordHandler(r.passwords, r.log)

	router.Route("/integrations/password", func(r chi.Router) {
		r.Get("/receiver/{address}", password.GetPasswordReceiver)
		r.Get("/sender/{address}", password.GetPasswordSender)
	})

	return router
}
