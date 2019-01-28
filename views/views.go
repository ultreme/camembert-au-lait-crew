package views

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/oxtoacart/bpool"

	"ultre.me/calcbiz/api"
	"ultre.me/calcbiz/pkg/crew"
)

type Options struct {
	Router       *mux.Router
	Debug        bool
	Svc          api.ServerServer
	StaticBox    *packr.Box
	TemplatesBox *packr.Box
}

type handlers struct {
	opts           *Options
	funcmap        *ctxFuncmap
	templatesMutex sync.Mutex
	templates      map[string]*template.Template
	bufpool        *bpool.BufferPool
}

func Setup(opts *Options) error {
	handlers := handlers{
		opts:    opts,
		bufpool: bpool.NewBufferPool(64),
	}
	if err := handlers.loadTemplates(); err != nil {
		return err
	}

	// FIXME: configure router to accept trailing slashes, but redirect

	// home
	opts.Router.HandleFunc("/", handlers.homeHandler)

	// muzik
	opts.Router.HandleFunc("/muzik", handlers.muzikHandler)
	opts.Router.HandleFunc("/track/{track_id:[0-9]+}", handlers.trackHandler)
	opts.Router.HandleFunc("/album/{album_id:[0-9]+}", handlers.albumHandler)

	// hackz
	opts.Router.HandleFunc("/hackz", handlers.hackzHandler)
	opts.Router.HandleFunc("/hackz/recettator", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/convertisseur", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/miroir", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/demineur", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/calculatrice.exe", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/terminul", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/steak-hache-shake", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/3615cryptage", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/paint", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/2048", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/ultreme-tetris", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/moijaime", handlers.hackzTravaux)
	opts.Router.HandleFunc("/hackz/phazms", handlers.hackzTravaux) // GET POST ?
	opts.Router.HandleFunc("/hackz/m1ch3l", handlers.hackzTravaux) // GET POST ?

	// copaings
	opts.Router.HandleFunc("/copaings", handlers.copaingsHandler)

	//
	// old routes (to be imported)
	//

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
	data := renderData{
		"track":       track,
		"layout_mode": "two_columns",
	}
	h.render(w, r, "track.tmpl", data)
}

func (h *handlers) albumHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumId, err := strconv.ParseUint(vars["album_id"], 10, 64)
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	album, err := h.opts.Svc.SoundcloudPlaylist(r.Context(), &api.SoundcloudPlaylistInput{
		PlaylistId: albumId,
	})
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	data := renderData{"album": album}
	h.render(w, r, "album.tmpl", data)
}

func (h *handlers) copaingsHandler(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	data := renderData{"friends": crew.CALC.Friends}
	h.render(w, r, "copaings.tmpl", data)
}

func (h *handlers) hackzTravaux(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.travaux.tmpl", nil)
}

func (h *handlers) hackzHandler(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	hackz, err := h.opts.Svc.Hackz(nil, &api.Void{})
	if err != nil {
		h.renderError(w, r, err)
		return
	}
	data := renderData{
		"hackz":       hackz,
		"layout_mode": "two_columns",
	}
	h.render(w, r, "hackz.tmpl", data)
}
