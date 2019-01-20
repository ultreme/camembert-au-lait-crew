package views

import (
	"html/template"
	"math/rand"
)

func funcmap(opts *Options) template.FuncMap {
	return template.FuncMap{
		"yomyman_style": yomymanStyle,
		"devel":         func() bool { return opts.Debug },
	}
}

func yomymanStyle() string {
	styles := []string{"cachou", "jambon", "epinard", "lasagne", "haricot", "sandwich"}
	return "cool-style-" + styles[rand.Intn(len(styles))]
}
