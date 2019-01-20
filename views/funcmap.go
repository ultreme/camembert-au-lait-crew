package views

import (
	"html/template"
	"math/rand"

	"go.uber.org/zap"
	"ultre.me/calcbiz/pkg/random"
)

func getFuncmap(opts *Options) template.FuncMap {
	f := funcmap{opts: opts}
	return template.FuncMap{
		"yomyman_style":           f.yomymanStyle,
		"devel":                   f.devel,
		"sharing_image_url":       f.sharingImageURL,
		"sharing_description":     f.sharingDescription,
		"current_url":             f.currentURL,
		"active_page":             f.activePage,
		"resize":                  f.resize,
		"page_title":              f.pageTitle,
		"invalid_cache":           f.invalidCache,
		"logo_alternate":          random.AlternateLogo,
		"mot_du_jour":             random.WOTD,
		"megahertz":               f.megahertz,
		"mot_debile_qui_se_mange": random.MotDebileQuiSeMange,
		"neige":                   func() bool { return false },
	}
}

type funcmap struct{ opts *Options }

func (f *funcmap) devel() bool { return f.opts.Debug }

func (f *funcmap) sharingImageURL() string {
	return "http://www.camembertaulaitcrew.biz/static/img/logo-300.png" // FIXME: make it dynamic
}

func (f *funcmap) activePage() string { return "home" } // FIXME: make it dynamic

func (f *funcmap) pageTitle() string { return "Camembert au lait crew" } // FIXME: make it dynamic

func (f *funcmap) sharingDescription() string { return "c'est cool" }

func (f *funcmap) currentURL() string {
	return "https://www.camembertaulaitcrew.biz/"
	// FIXME: make it flexible (should be canonical url
}

func (f *funcmap) resize(opts ...string) string {
	path := opts[len(opts)-1]
	opts = opts[:len(opts)-1]
	zap.L().Debug("redize", zap.String("path", path), zap.Strings("opts", opts))
	// FIXME: apply transform + cache + give new URL
	return path
}

func (f *funcmap) yomymanStyle() string {
	styles := []string{"cachou", "jambon", "epinard", "lasagne", "haricot", "sandwich"}
	return "cool-style-" + styles[rand.Intn(len(styles))]
}

func (f *funcmap) invalidCache() string {
	// FIXME: use random string (once per server start)
	// FIXME: or create a new func that returns file content hash
	return "no_cache_please"
}

func (f *funcmap) megahertz() float64 {
	return rand.Float64() * 100
}
