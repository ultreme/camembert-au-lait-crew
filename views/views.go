package views

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) error {
	if err := loadTemplates(); err != nil {
		return err
	}
	router.HandleFunc("/", homeHandler)
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	data := map[string]interface{}{
		"hello": "world",
	}
	render(w, r, "home.tmpl", data)
}
