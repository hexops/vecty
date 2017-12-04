package asset

import (
	"net/http"

	"github.com/gopherjs/vecty/internal/gopherjs_http"
	"github.com/nytimes/gziphandler"
)

// Options defines options for the asset server.
type Options struct {
	Dir         string
	StripPrefix string
	Gzip        bool
}

// Defaults returns the default options for the asset server:
//
// 	&Options{
// 		Dir:         "assets",
// 		StripPrefix: "/assets/",
// 		Gzip:        true,
// 	}
//
func Defaults() *Options {
	return &Options{
		Dir:         "assets",
		StripPrefix: "/assets/",
		Gzip:        true,
	}
}

// NewServer returns a new asset server with the given options. If o == nil,
// then Defaults() are used as the options.
func NewServer(o *Options) http.Handler {
	if o == nil {
		o = Defaults()
	}

	assets := gopherjs_http.NewFS(http.Dir(o.Dir))

	h := http.StripPrefix(o.StripPrefix, http.FileServer(assets))
	if o.Gzip {
		h = gziphandler.GzipHandler(h)
	}
	return h
}
