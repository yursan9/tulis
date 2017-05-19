package post

import (
	"log"
	"path/filepath"
	"sort"
)

//GetPosts take path to directory for the post
func GetPosts(postDir string) []*Post {
	//Glob all the files in postDir that end with '.md'
	paths, err := filepath.Glob(filepath.Join(postDir, "*.md"))
	if err != nil {
		log.Print(err)
	}

	all := make([]*Post, 0, len(paths))
	for _, path := range paths {
		p, err := New(path)
		if err != nil {
			log.Print(err)
		}

		all = append(all, p)
	}
	// Sort the Post in descending order
	sort.Slice(all, func(i, j int) bool {
		return all[i].Meta.Date.Time.After(all[j].Meta.Date.Time)
	})
	return all
}
