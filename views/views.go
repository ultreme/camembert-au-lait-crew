package views

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"ultre.me/calcbiz/api"
	"ultre.me/calcbiz/pkg/crew"
)

type Options struct {
	Router *mux.Router
	Debug  bool
	Svc    api.ServerServer
}

func Setup(opts *Options) error {
	handlers := handlers{opts: opts}
	if err := handlers.loadTemplates(); err != nil {
		return err
	}
	opts.Router.HandleFunc("/", handlers.homeHandler)
	opts.Router.HandleFunc("/muzik", handlers.muzikHandler)
	opts.Router.HandleFunc("/hackz", handlers.hackzHandler)
	opts.Router.HandleFunc("/copaings", handlers.copaingsHandler)
	opts.Router.HandleFunc("/track/{track_id:[0-9]+}", handlers.trackHandler)

	//
	// old routes (to be imported)
	//

	// /m1ch3l GET/POST
	// /hackz/convertisseur
	// /hackz/miroir
	// /hackz/demineur
	// /hackz/calculatrice.exe
	// /hackz/terminul
	// /hackz/steak-hache-shake
	// /hackz/3615cryptage
	// /hackz/paint
	// /hackz/2048
	// /hackz/ultreme-tetris
	// /hackz/recettator
	// /hackz/moijaime
	// /hackz/phazms GET/POST
	// /album/<int:album_id>
	// /scorz/inc/<string:user>/<string:what>/<int:points>
	// /sitemap.xml

	//
	// old inactive routes
	//

	// /admin/flush-cache
	// /vidz
	// /tofz
	// /test/fb
	// /test/tototo
	// /scorz

	//
	// new routes
	//

	// /crew
	return nil
}

type handlers struct {
	opts    *Options
	funcmap *ctxFuncmap
	mutex   sync.Mutex
}

type renderData map[string]interface{}

func (h *handlers) homeHandler(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	dashboard, err := h.opts.Svc.Dashboard(nil, &api.Void{})
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	data := renderData{"dashboard": dashboard}
	h.render(w, r, "home.tmpl", data)
}

func (h *handlers) muzikHandler(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	playlists, err := h.opts.Svc.SoundcloudPlaylists(r.Context(), &api.Void{})
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	data := renderData{
		"playlists": playlists,
	}
	h.render(w, r, "muzik.tmpl", data)
}

func (h *handlers) trackHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackId, err := strconv.ParseUint(vars["track_id"], 10, 64)
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	track, err := h.opts.Svc.SoundcloudTrack(r.Context(), &api.SoundcloudTrackInput{
		TrackId: trackId,
	})
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	data := renderData{"track": track}
	h.render(w, r, "track.tmpl", data)
}

func (h *handlers) copaingsHandler(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	data := renderData{"friends": crew.CALC.Friends}
	h.render(w, r, "copaings.tmpl", data)
}

func (h *handlers) hackzHandler(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	hackz, err := h.opts.Svc.Hackz(nil, &api.Void{})
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	data := renderData{"hackz": hackz}
	h.render(w, r, "hackz.tmpl", data)
}
