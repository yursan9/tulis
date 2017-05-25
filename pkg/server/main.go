package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/yursan9/tulis/pkg/post"
)

var (
	port        string
	templateDir string
	postDir     string
	staticDir   string
	base        string
	rel         bool
	all         []*post.Post
	maxPost     uint8
)

func init() {
	// Declare ok first
	ok := true
	// Use only assignment, because using shorthand shadow global port variabel
	port, ok = os.LookupEnv("TULIS_PORT")
	if !ok {
		port = ":8080"
	}

	templateDir, ok = os.LookupEnv("TULIS_TEMPLDIR")
	if !ok {
		templateDir = "templates"
	}
	postDir, ok = os.LookupEnv("TULIS_POSTDIR")
	if !ok {
		postDir = "posts"
	}
	staticDir, ok = os.LookupEnv("TULIS_STATICDIR")
	if !ok {
		staticDir = "static"
	}

	var err error
	rel, err = strconv.ParseBool(os.Getenv("TULIS_RELATIVE"))
	if err != nil {
		rel = false
	}
	if rel {
		base, ok = os.LookupEnv("TULIS_BASE")
		if !ok {
			base = os.Getenv("PWD")
		}
		templateDir = filepath.Join(base, templateDir)
		postDir = filepath.Join(base, postDir)
		staticDir = filepath.Join(base, staticDir)
	}

	if n, err := strconv.ParseUint(os.Getenv("TULIS_MAXPOSTS"), 10, 8); err != nil {
		maxPost = 5
	} else {
		maxPost = uint8(n)
	}

	all = post.GetPosts(postDir)
}

func newRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/page/:page", Page)
	r.GET("/post/:title", ReadPost)
	r.GET("/about", About)
	r.GET("/tag/:name", ByTag)
	r.ServeFiles("/static/*filepath", http.Dir(staticDir))

	return r
}

// Make new server
func New() *http.Server {
	r := newRouter()
	return &http.Server{
		Addr:    port,
		Handler: r,
	}
}
