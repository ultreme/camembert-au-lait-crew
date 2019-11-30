package views

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/oxtoacart/bpool"
	"ultre.me/calcbiz/pkg/api"
	"ultre.me/calcbiz/pkg/crew"
	"ultre.me/recettator"
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
	opts.Router.HandleFunc("/hackz/recettator", handlers.hackzRecettator)
	opts.Router.HandleFunc("/hackz/convertisseur", handlers.hackzConvertisseur)
	opts.Router.HandleFunc("/hackz/miroir", handlers.hackzMiroir)
	opts.Router.HandleFunc("/hackz/demineur", handlers.hackzDemineur)
	opts.Router.HandleFunc("/hackz/calculatrice.exe", handlers.hackzCalculatrice)
	opts.Router.HandleFunc("/hackz/terminul", handlers.hackzTerminul)
	opts.Router.HandleFunc("/hackz/steak-hache-shake", handlers.hackzSteakHacheShake)
	opts.Router.HandleFunc("/hackz/3615cryptage", handlers.hackz3615cryptage)
	opts.Router.HandleFunc("/hackz/paint", handlers.hackzPaint)
	opts.Router.HandleFunc("/hackz/2048", handlers.hackz2048)
	opts.Router.Path("/hackz/ultreme-tetris").HandlerFunc(handlers.hackzTetris).Name("hackz.tetris")
	opts.Router.HandleFunc("/hackz/moijaime", handlers.hackzMoiJaime)
	opts.Router.HandleFunc("/hackz/phazms", handlers.hackzPhazms) // GET POST ?
	opts.Router.HandleFunc("/hackz/m1ch3l", handlers.hackzM1ch3l) // GET POST ?

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

func (h *handlers) hackzCalculatrice(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.calculatrice.tmpl", nil)
}

func (h *handlers) hackz2048(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	data := renderData{
		"layout_mode": "two_columns",
	}
	h.render(w, r, "hackz.2048.tmpl", data)
}

func (h *handlers) hackzSteakHacheShake(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.steakhacheshake.tmpl", nil)
}

func (h *handlers) hackzConvertisseur(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.convertisseur.tmpl", nil)
}

func (h *handlers) hackzMiroir(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.miroir.tmpl", nil)
}

func (h *handlers) hackzTerminul(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.terminul.tmpl", nil)
}

func (h *handlers) hackzDemineur(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.demineur.tmpl", nil)
}

func (h *handlers) hackzTetris(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	niveau := r.FormValue("niveau")
	if niveau == "" {
		niveau = "5"
	}
	data := renderData{
		"layout_mode": "two_columns",
		"niveau":      niveau,
	}

	h.render(w, r, "hackz.tetris.tmpl", data)
}

func (h *handlers) hackzMoiJaime(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.moijaime.tmpl", nil)
}

func (h *handlers) hackzPaint(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.paint.tmpl", nil)
}

func (h *handlers) hackzRecettator(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)

	seed := r.FormValue("recette")
	if seed == "" {
		http.Redirect(w, r, fmt.Sprintf("/hackz/recettator?recette=%d", rand.Intn(10000+1)), http.StatusFound)
		return
	}

	seedInt, err := strconv.Atoi(seed)
	if err != nil {
		log.Printf("invalid seed: %v (err=$v)", seed, err)
		seedInt = 42
	}

	rctt := recettator.New(int64(seedInt))
	rctt.SetSettings(recettator.Settings{
		//MainIngredients:      uint64(rand.Intn(4) + 1),
		//SecondaryIngredients: uint64(rand.Intn(4) + 1),
		//Steps:                uint64(rand.Intn(5) + 2),
		// People: 100
	})

	otherRecettes := []string{}
	for i := 0; i < 10; i++ {
		newSeed := (rand.Intn(588-59) + 59) * 17 // limits to 529 results between 1000 and 10000 (for SEO)
		otherRecettes = append(otherRecettes, fmt.Sprintf("%d", newSeed))
	}

	data := renderData{
		"layout_mode":   "two_columns",
		"seed":          seed,
		"seedInt":       seedInt,
		"recettator":    rctt,
		"Title":         rctt.Title(),
		"People":        rctt.People(),
		"Steps":         rctt.Steps(),
		"Pool":          rctt.Pool(),
		"otherRecettes": otherRecettes,
	}
	h.render(w, r, "hackz.recettator.tmpl", data)
}

func (h *handlers) hackzPhazms(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.phazms.tmpl", nil)
}

func (h *handlers) hackzM1ch3l(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.m1ch3l.tmpl", nil)
}

func (h *handlers) hackz3615cryptage(w http.ResponseWriter, r *http.Request) {
	h.setDefaultHeaders(w)
	h.render(w, r, "hackz.3615cryptage.tmpl", nil)
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
