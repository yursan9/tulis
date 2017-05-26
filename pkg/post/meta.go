package post

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gosimple/slug"
	"github.com/russross/blackfriday"
)

var (
	errInvalidFrontMatter = errors.New("Invalid front matter")
	errMissingFrontMatter = errors.New("Missing front matter")

	dateFmt = map[int]string{
		10: "2006-01-02",
		15: "2006-01-02 15:04",
		19: "2006-01-02 15:04 WIB",
		25: time.RFC3339,
	}
)

type date struct {
	time.Time
}

func (d *date) UnmarshalText(t []byte) error {
	var err error
	d.Time, err = time.Parse(dateFmt[len(string(t))], string(t))
	return err
}

// Meta structure contain basic metadata for blog post
type Meta struct {
	Title       string
	Author      string
	Description string
	Date        date
	Tag         []string
}

// Post structure combine the metadata and content
type Post struct {
	*Meta
	Slug    string
	Content template.HTML
}

func newMeta(t string) (*Meta, error) {
	m := new(Meta)
	if _, err := toml.Decode(t, &m); err != nil {
		log.Fatal(err)
	}
	return m, nil
}

//New generate Post from fn
//fn must be valid filepath
func New(fn string) (*Post, error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	//Check if document start with valid token
	if !bytes.HasPrefix(b, []byte("---\n")) {
		return nil, errMissingFrontMatter
	}
	b = bytes.TrimPrefix(b, []byte("---\n"))

	//Split b to array, array[0] is front matter
	//array[1] is the rest of text (post's body)
	arr := bytes.SplitN(b, []byte("\n---\n"), 2)

	//Generate meta from text
	m, err := newMeta(string(arr[0]))

	//Convert the rest of text to Markdown
	body := blackfriday.MarkdownCommon(arr[1])
	p := &Post{
		m,
		slug.Make(m.Title),
		template.HTML(body),
	}
	return p, nil
}
