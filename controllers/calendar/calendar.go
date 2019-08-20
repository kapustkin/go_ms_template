package calendar

import (
	"net/http"

	"github.com/go-chi/chi"
)

// GetUserCalandar return calendar for user
func GetUserCalandar(res http.ResponseWriter, req *http.Request) {
	user := chi.URLParam(req, "user")
	//nolint:errcheck
	res.Write([]byte("Hello " + user))
}
