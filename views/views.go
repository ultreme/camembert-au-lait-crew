package views

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	router.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	data := map[string]interface{}{
		"hello": "world",
	}
	render(w, r, "home.html", data)
}
