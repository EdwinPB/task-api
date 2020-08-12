package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:         "application.plush.html",
		DefaultContentType: "application/json",
		TemplatesBox:       packr.New("app:templates", "../templates"),
	})
}
