package asset_test

import (
	"log"
	"net/http"

	"github.com/gopherjs/vecty/asset"
)

func ExampleAssetServer() {
	http.Handle("/assets", asset.NewServer(nil))
	log.Fatal(http.ListenAndServe(":8777", nil))
}
