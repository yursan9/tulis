package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index handle request for our index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := newPageData(1)
	if err != nil {
		http.NotFound(w, r)
	}
	t["index"].Execute(w, data)
}

// ReadPost handle request for reading our posts
func ReadPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := newPostData(ps.ByName("title"))
	if err != nil {
		http.NotFound(w, r)
	}
	t["post"].Execute(w, data)
}

// Page handle per page view, actually index is an alias to page 1
func Page(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Halaman: %d", ps.ByName("page"))
}

// About handle request to serve an about page
func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "About Me!")
}

// ByTag handle request to search posts based on its tag.
func ByTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Search by tag: %s", ps.ByName("name"))
}
