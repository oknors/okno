package mod

import (
	"html/template"
	"time"

	"github.com/russross/blackfriday/v2"
	stripmd "github.com/writeas/go-strip-markdown"
)

// Post contains articles and pages used by the CMS
type Post struct {
	ID             string `schema:"id,required"`
	Title          string
	Taxonomies     map[string]Taxonomy
	Content        string
	ContentPreview string
	CustomFields   map[string]interface{}
	Slug           string
	IsDraft        bool
	Published      bool
	Active         bool
	ItemType       string
	Template       string
	Order          int
	Author         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Taxonomy struct {
	ID             string `schema:"id,required"`
	Title          string
	Taxonomies     map[string]Taxonomy
	Content        string
	ContentPreview string
	CustomFields   map[string]interface{}
	Slug           string
}

// GetHTMLContent returns the post's markdown content as HTML
func (p *Post) GetHTMLContent() template.HTML {
	str := string(blackfriday.Run([]byte(p.Content)))
	return template.HTML(str)
}

// GetContentPreview strips markdown from the content string, trims and returns it
func (p *Post) GetContentPreview(max int) string {
	preview := stripmd.Strip(string(p.Content))
	if len(preview) > max {
		preview = preview[:max]
	}
	return preview
}

//func (posts Posts)Sort(){
//	for _, p := range posts {
//		posts[p.Order] = p
//		fmt.Println("Initial:",p.Order)
//	}
//	sort.Sort(sort.Reverse(posts))
//	return
//}

type Posts []Post

func (p Posts) Len() int           { return len(p) }
func (p Posts) Less(i, j int) bool { return p[i].Order < p[j].Order }
func (p Posts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
