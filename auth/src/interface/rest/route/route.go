package route

import (
	userHandler "efishery/auth/src/interface/rest/handlers/users"
	"efishery/auth/src/interface/rest/middleware"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserAppRouter(ch userHandler.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", UserRouter(ch))

	return r
}
