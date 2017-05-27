package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Index handle request for our index page, it's an alias for Page 1
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := newPageData(1)
	if err != nil {
		http.NotFound(w, r)
	}
	err = t["index"].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadPost handle request for reading our posts
func ReadPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := newPostData(ps.ByName("title"))
	if err != nil {
		http.NotFound(w, r)
	}
	err = t["post"].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

// Page handle per page view, actually index is an alias to page 1
func Page(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p, err := strconv.ParseUint(ps.ByName("page"), 10, 8)
	if err != nil {
		http.NotFound(w, r)
	}
	data, err := newPageData(uint8(p))
	if err != nil {
		http.NotFound(w, r)
	}
	err = t["index"].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

// About handle request to serve an about page
func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "About Me!")
}

// ByTag handle request to search posts based on its tag.
func ByTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p, err := strconv.ParseUint(ps.ByName("page"), 10, 8)
	if err != nil {
		http.NotFound(w, r)
	}
	data := newByTagData(uint8(p), ps.ByName("name"))
	if err != nil {
		http.NotFound(w, r)
	}
	err = t["tag"].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

