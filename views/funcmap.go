package views

import (
	"fmt"
	"html/template"
	"image"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/moul/sprig"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"

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
	fm["mot_cool"] = random.MotCool
	fm["megahertz"] = f.megahertz
	fm["mot_debile_qui_se_mange"] = random.MotDebileQuiSeMange
	fm["neige"] = func() bool { return false }
	fm["cache_external_asset"] = f.cacheExternalAsset
	fm["linkify"] = f.linkify
	f.fm = fm
	return f
}

type ctxFuncmap struct {
	fm   template.FuncMap
	opts *Options
	req  *http.Request
}

func (f *ctxFuncmap) linkify(input string) string {
	// FIXME: replace URLs with links
	return input
}

func (f *ctxFuncmap) devel() bool { return f.opts.Debug }

func (f *ctxFuncmap) sharingImageURL() string {
	return "http://www.camembertaulaitcrew.biz/static/img/logo-300.png" // FIXME: make it dynamic
}

func (f *ctxFuncmap) activePage() string {
	return f.req.URL.Path
}

func (f *ctxFuncmap) activeMenu() string {
	switch path := f.req.URL.Path; {
	case path == "/":
		return "home"
	case path == "/muzik" ||
		strings.HasPrefix(f.req.URL.Path, "/track/") ||
		strings.HasPrefix(f.req.URL.Path, "/album/"):
		return "muzik"
	case path == "/copaings":
		return "copaings"
	case path == "/hackz" ||
		strings.HasPrefix(f.req.URL.Path, "/hackz"):
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

func (f *ctxFuncmap) cacheExternalAsset(input string) (string, error) {
	if !strings.HasPrefix(input, "http") {
		return input, nil
	}

	h := make([]byte, 8)
	sha3.ShakeSum256(h, []byte(input))
	// fixme parse url
	newpath := fmt.Sprintf("./static/img/cache/%x%s", h, filepath.Ext(input))

	out, err := os.Create(newpath)
	if err != nil {
		return input, err
	}
	defer out.Close()
	resp, err := http.Get(input)
	if err != nil {
		return input, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return input, err
	}

	return newpath, nil
}

func (f *ctxFuncmap) resize(opts ...string) string {
	path := opts[len(opts)-1]
	opts = opts[:len(opts)-1]
	urlAppend := ""
	if f.opts.Debug {
		urlAppend += "?src=" + path
	}

	var err error
	if path, err = f.cacheExternalAsset(path); err != nil {
		zap.L().Warn("failed to get external asset", zap.String("path", path), zap.Error(err))
		return path
	}

	buf := []byte(fmt.Sprintf("%s:%v", path, opts))
	h := make([]byte, 8)
	sha3.ShakeSum256(h, buf)
	//FIXME: process hash based on file content instead of filepath (keep opts)
	newpath := fmt.Sprintf("./static/img/cache/%x%s", h, filepath.Ext(path))

	if _, err := os.Stat(newpath); !os.IsNotExist(err) {
		return strings.Replace(newpath, "./static/", "/", -1) + urlAppend
	}

	logger := zap.L().With(
		zap.String("srcpath", path),
		zap.String("destpath", newpath),
		zap.Strings("opts", opts),
		zap.String("buf", string(buf)),
		zap.String("hash", fmt.Sprintf("%x", h)),
	)

	src, err := imaging.Open(path)
	if err != nil {
		logger.Warn("failed to open image", zap.Error(err))
		return path
	}

	var newimg image.Image
	for _, opt := range opts {
		spl := strings.Split(opt, "=")
		if len(spl) != 2 {
			logger.Warn("invalid options", zap.String("opt", opt))
			return path
		}
		switch spl[0] {
		case "fill":
			dims := strings.Split(spl[1], "x")
			if len(dims) != 2 {
				logger.Warn("invalid dimensions", zap.String("opt", opt))
				return path
			}
			width, err1 := strconv.ParseInt(dims[0], 10, 64)
			height, err2 := strconv.ParseInt(dims[1], 10, 64)
			if err1 != nil || err2 != nil {
				logger.Warn("invalid dimensions (not a number)", zap.String("opt", opt))
				return path
			}
			newimg = imaging.Fill(src, int(width), int(height), imaging.Center, imaging.Lanczos)
		default:
			logger.Warn("unhandled option type", zap.String("opt", opt))
			return path
		}
	}

	if newimg == nil {
		logger.Warn("no operation done on image, doing nothing")
		return path
	}

	if err := imaging.Save(newimg, newpath); err != nil {
		logger.Warn("failed to save resized image", zap.Error(err))
		return path
	}
	return strings.Replace(newpath, "./static/", "/", -1) + urlAppend
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
