package views

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/moul/sprig"
	"ultre.me/calcbiz/pkg/random"
)

func getFuncmap(opts *Options) *ctxFuncmap {
	f := &ctxFuncmap{opts: opts}

	fm := sprig.FuncMap()
	fm["yomyman_style"] = f.yomymanStyle
	fm["devel"] = f.devel
	fm["sharing_image_url"] = f.sharingImageURL
	fm["sharing_description"] = f.sharingDescription
	fm["current_url"] = f.currentURL
	fm["active_page"] = f.activePage
	fm["active_menu"] = f.activeMenu
	fm["resize"] = f.resize
	fm["page_title"] = f.pageTitle
	fm["invalid_cache"] = f.invalidCache
	fm["logo_alternate"] = random.AlternateLogo
	fm["mot_du_jour"] = random.WOTD
	fm["megahertz"] = f.megahertz
	fm["mot_debile_qui_se_mange"] = random.MotDebileQuiSeMange
	fm["neige"] = func() bool { return false }
	fm["cache_external_assets"] = f.cacheExternalAsset
	f.fm = fm
	return f
}

type ctxFuncmap struct {
	fm   template.FuncMap
	opts *Options
	req  *http.Request
}

func (f *ctxFuncmap) devel() bool { return f.opts.Debug }

func (f *ctxFuncmap) sharingImageURL() string {
	return "http://www.camembertaulaitcrew.biz/static/img/logo-300.png" // FIXME: make it dynamic
}

func (f *ctxFuncmap) activePage() string {
	return f.req.URL.Path
}

func (f *ctxFuncmap) activeMenu() string {
	switch f.req.URL.Path {
	case "/":
		return "home"
	case "/muzik":
		// FIXME: support albums, songs
		return "muzik"
	case "/copaings":
		return "copaings"
	case "/hackz":
		// FIXME: support hackz URLs
		return "hackz"
	default:
		return "home"
	}
}

func (f *ctxFuncmap) pageTitle() string { return "Camembert au lait crew" } // FIXME: make it dynamic

func (f *ctxFuncmap) sharingDescription() string { return "c'est cool" }

func (f *ctxFuncmap) currentURL() string {
	return "https://www.camembertaulaitcrew.biz/"
	// FIXME: make it flexible (should be canonical url
}

func (f *ctxFuncmap) cacheExternalAsset(path string) string {
	// FIXME: implement
	return path
}

func (f *ctxFuncmap) resize(opts ...string) string {
	path := opts[len(opts)-1]
	opts = opts[:len(opts)-1]
	//zap.L().Debug("resize", zap.String("path", path), zap.Strings("opts", opts))
	// FIXME: apply transform + cache + give new URL
	return path
}

func (f *ctxFuncmap) yomymanStyle() string {
	styles := []string{"cachou", "jambon", "epinard", "lasagne", "haricot", "sandwich"}
	return "cool-style-" + styles[rand.Intn(len(styles))]
}

func (f *ctxFuncmap) invalidCache() string {
	// FIXME: use random string (once per server start)
	// FIXME: or create a new func that returns file content hash
	return "no_cache_please"
}

func (f *ctxFuncmap) megahertz() float64 {
	return rand.Float64() * 100
}
