package server

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/yursan9/tulis/pkg/post"
)

var (
	t        *template.Template
	funcMaps template.FuncMap
)

func init() {
	// Define custom filter for template
	funcMaps = template.FuncMap{
		"inc": func(i uint8) uint8 { return i + 1 },
		"dec": func(i uint8) uint8 { return i - 1 },
	}
	// Initialize map of template
	t = template.New("").Funcs(funcMaps)
	template.Must(t.ParseGlob(filepath.Join(templateDir, "*")))
}

// PageData contain struct for Page and Index template
type PageData struct {
	PageNow uint8
	PageMax uint8
	Posts   []*post.Post
}

func newPageData(n uint8) (*PageData, error) {
	pd := new(PageData)

	// Initialize page number
	pd.PageMax = uint8(len(all))/maxPost + 1
	if n < 1 || n > pd.PageMax {
		return nil, fmt.Errorf("There is no page %d", n)
	}
	pd.PageNow = n

	// Initialize array of posts
	s := maxPost * (n - 1)
	pd.Posts = all[s:]
	f := maxPost * n
	if f < uint8(len(all)) {
		pd.Posts = all[s:f]
	}

	return pd, nil
}

// PostData contain struct for Post template
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

// ByTagData contain struct for Search By Tag template
type ByTagData struct {
	PageNow uint8
	PageMax uint8
	Tag     string
	Posts   []*post.Post
}

func newByTagData(n uint8, tag string) *ByTagData {
	pd := new(ByTagData)
	pd.Tag = tag
	// Search post with given tag
OUTER:
	for _, p := range all {
		for _, t := range p.Tag {
			if tag == t {
				pd.Posts = append(pd.Posts, p)
				continue OUTER
			}
		}
	}
	if len(pd.Posts) == 0 {
		pd.PageNow = 1
		pd.PageMax = 1
		return pd
	}

	// Initialize Page number
	pd.PageNow = n
	pd.PageMax = uint8(len(pd.Posts))/maxPost + 1

	// Initialize array of posts
	s := maxPost * (n - 1)
	pd.Posts = pd.Posts[s:]
	f := maxPost * n
	if f < uint8(len(pd.Posts)) {
		pd.Posts = pd.Posts[s:f]
	}

	return pd
}

