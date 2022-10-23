package route

import (
	"net/http"

	userHandler "efishery/auth/src/interface/rest/handlers/users"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h userHandler.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.Post("/login", h.CreateToken)
	r.Post("/info", h.InfoToken)
	return r
}
