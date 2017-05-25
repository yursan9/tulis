package server

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/yursan9/tulis/pkg/post"
)

var (
	t map[string]*template.Template
)

func init() {
	t = map[string]*template.Template{
		"index": template.Must(template.ParseFiles(filepath.Join(
			templateDir, "index.html"))),
		"post": template.Must(template.ParseFiles(filepath.Join(
			templateDir, "post.html"))),
	}
}

type PageData struct {
	PageNow  uint8
	PageNext uint8
	PagePrev uint8
	Posts    []*post.Post
}

func newPageData(n uint8) (*PageData, error) {
	maxPage := uint8(len(all))/maxPost + 1
	if n < 1 || n > maxPage {
		return nil, fmt.Errorf("There is no page %d", n)
	}

	pd := new(PageData)
	// Initialize page number
	pd.PageNow = n
	if n > 1 {
		pd.PagePrev = n - 1
	}
	if n < maxPage {
		pd.PageNext = n + 1
	}
	
	// Initialize array of posts
	s := maxPost * (n - 1)
	pd.Posts = all[s:]
	f := maxPost * n
	if f < uint8(len(all)) {
		pd.Posts = all[s:f]
	}

	return pd, nil
}

type PostData struct {
	*post.Post
	Next *post.Post
	Prev *post.Post
}

func newPostData(slug string) (*PostData, error) {
	pd := new(PostData)
	for i, p := range all {
		if p.Slug == slug {
			pd.Post = p
			if i > 0 {
				pd.Prev = all[i-1]
			}
			if i < len(all)-1 {
				pd.Next = all[i+1]
			}
			return pd, nil
		}
	}
	return nil, fmt.Errorf("Can't find post with given slug: %s", slug)
}

