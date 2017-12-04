// The asset package provides an asset serving HTTP handler with live GopherJS code
// recompilation.
//
// BUGS:
//
// - How do we handle embedded assets ala vfsgendev etc?
// - Source maps do not work
// - Maybe asset.NewServer is not the best name?
// - Maybe asset.NewServer should take a dir argument AND options, not just options.
// - Should github.com/nytimes/gziphandler be internal ? probably
// - How fast is internal/gopherjs_http ? should there be a production mode where no code recompilation occurs?
// -
//
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
