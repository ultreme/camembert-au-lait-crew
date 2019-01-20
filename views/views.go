package views

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Options struct {
	Router *mux.Router
	Debug  bool
}

func Setup(opts *Options) error {
	if err := loadTemplates(opts); err != nil {
		return err
	}
	opts.Router.HandleFunc("/", homeHandler)
	opts.Router.HandleFunc("/muzik", muzikHandler)
	opts.Router.HandleFunc("/hackz", hackzHandler)
	opts.Router.HandleFunc("/copaings", copaingsHandler)
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	data := map[string]interface{}{"hello": "world"}
	render(w, r, "home.tmpl", data)
}

func muzikHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	data := map[string]interface{}{"hello": "world"}
	render(w, r, "muzik.tmpl", data)
}

func copaingsHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	data := map[string]interface{}{"hello": "world"}
	render(w, r, "copaings.tmpl", data)
}

func hackzHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	data := map[string]interface{}{"hello": "world"}
	render(w, r, "hackz.tmpl", data)
}
