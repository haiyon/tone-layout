package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

// NewCORS - Cross Origin Resource Sharing
func NewCORS() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.ExposedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.MaxAge(int(10*time.Minute/time.Second)),
		handlers.OptionStatusCode(http.StatusMisdirectedRequest),
	)
}
