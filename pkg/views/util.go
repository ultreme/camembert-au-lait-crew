package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"

	"github.com/gobuffalo/packd"
	"go.uber.org/zap"
)

func (h *handlers) loadTemplates() error {
	h.templates = make(map[string]*template.Template)

	// load template files
	layoutContent := ""
	pageContents := map[string]string{}
	err := h.opts.TemplatesBox.Walk(func(filepath string, file packd.File) error {
		if strings.HasPrefix(path.Base(filepath), ".#") {
			// ignore temporary files
			return nil
		}
		switch path.Dir(filepath) {
		case ".":
			pageContents[filepath] = file.String()
		case "layout":
			layoutContent += file.String()
		}
		return nil
	})
	if err != nil {
		return err
	}

	// generate optimized templates
	h.funcmap = getFuncmap(h.opts)
	mainTemplate := template.New("main").Funcs(h.funcmap.fm)
	mainTemplate = template.Must(mainTemplate.Parse(`{{define "main"}}{{template "base" .}}{{end}}`))
	mainTemplate = template.Must(mainTemplate.Parse(layoutContent))
	for filepath, content := range pageContents {
		h.templates[filepath] = template.Must(mainTemplate.Clone())
		h.templates[filepath] = template.Must(h.templates[filepath].Parse(content))
	}
	zap.L().Debug("templates loaded")
	return nil
}

func (h *handlers) setDefaultHeaders(w http.ResponseWriter) {
	h.push(w, "/css/calc.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (h *handlers) renderError(w http.ResponseWriter, r *http.Request, err error) {
	zap.L().Warn("rendering error", zap.Error(err))
	// FIXME: log in analytics
	if h.opts.Debug {
		http.Error(w, fmt.Sprintf("Error: %+v\n", err), http.StatusInternalServerError)
	} else {
		// FIXME: use error template
		http.Error(w, "Internal server error\n", http.StatusInternalServerError)
	}
}

func (h *handlers) render(w http.ResponseWriter, r *http.Request, name string, data renderData) {
	if h.opts.Debug {
		if err := h.loadTemplates(); err != nil {
			h.renderError(w, r, err)
			return
		}
	}

	tmpl, ok := h.templates[name]
	if !ok {
		h.renderError(w, r, fmt.Errorf("the template %s does not exist.", name))
		return
	}

	buf := h.bufpool.Get()
	defer h.bufpool.Put(buf)

	if data == nil {
		data = make(renderData)
	}

	data["req"] = map[string]interface{}{
		"name": name,
		"r":    r,
	}
	if _, found := data["layout_mode"]; !found {
		data["layout_mode"] = "single_column"
	}

	// set current request in ctxFuncmap objects
	h.templatesMutex.Lock()
	defer h.templatesMutex.Unlock()
	h.funcmap.req = r

	if err := tmpl.Execute(buf, data); err != nil {
		h.renderError(w, r, err)
		return
	}

	if _, err := buf.WriteTo(w); err != nil {
		zap.L().Warn("buf.WriteTo", zap.Error(err))
	}
}

func (h *handlers) push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			zap.L().Warn("failed to push resource", zap.String("path", resource), zap.Error(err))
			return
		}
	}
}
