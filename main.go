package main

import (
	"fmt"
	"github.com/yursan9/tulis/pkg/post"
)

func main() {
	p := post.GetPosts("posts")
	fmt.Println(p)
}
