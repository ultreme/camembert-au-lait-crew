package views

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gobuffalo/packr"
	"go.uber.org/zap"
)

var box = packr.NewBox("../templates")

func setDefaultHeaders(w http.ResponseWriter) {
	// push(w, "/static/xxx.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func renderError(w http.ResponseWriter, r *http.Request, err error) {
	zap.L().Warn("rendering error", zap.Error(err))
	fmt.Fprintf(w, "Error: %v\n", err)
}

func render(w http.ResponseWriter, r *http.Request, tplPath string, data interface{}) {
	tplFile, err := box.FindString(tplPath)
	if err != nil {
		renderError(w, r, err)
		return
	}
	tpl := template.Must(template.New(tplPath).Parse(tplFile))

	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, tplPath, data); err != nil {
		renderError(w, r, err)
		return
	}
	w.Write(buf.Bytes())
}
