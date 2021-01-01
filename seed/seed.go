package seed

import (
	"github.com/p9c/okno/app/jdb"
	"github.com/p9c/okno/app/models/post"
)

type postImport struct {
	ID        string
	Title     string
	Content   string
	Slug      string
	CreatedAt string
	UpdatedAt string
}

var (
	timeLayout = "22. January 2014."
)

func SeedOKNO(j *jdb.JDB) error {
	j.Collection("posts")
	err := seedPosts(j, getOKNOPosts)
	if err != nil {
		return err
	}
	return nil
}

func SeedParallelCoin(j *jdb.JDB) error {
	j.Collection("pages")
	err := seedPosts(j, getParallelCoinPages)
	if err != nil {
		return err
	}
	return nil
}

func seedPosts(j *jdb.JDB, s func() []post.Post) error {
	posts := []post.Post{}
	for _, p := range s() {
		post, err := post.Create(
			j,
			p.Title,
			p.Content,
			p.Slug,
			p.IsDraft,
		)
		if err != nil {
			return err
		}
		posts = append(posts, post)
	}
	return nil
}
