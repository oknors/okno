package seed

import (
	"github.com/p9c/okno/app/models/post"
)

func getParallelCoinPages() (pages []post.Post) {
	for _, p := range pagesParallelCoin {
		pages = append(pages, post.Post{
			Title:     p.Title,
			Content:   p.Content,
			Slug:      p.Slug,
			IsDraft:   false,
			Published: true,
			ItemType:  "page",
			Template:  "default",
		})
	}

	return
}

var (
	pagesParallelCoin = []post.Post{
		post.Post{
			ID:      "intro",
			Title:   "Plan 9 from CryptoSpace?",
			Content: "Let's Go Parallel",
			Slug:    "intro",
		},
		post.Post{
			ID:      "what",
			Title:   "What is ParallelCoin?",
			Content: `<p>The first question everyone asks is, “So, tell me, what is it you are doing and why is it unique and better and why should I continue to look into this?”</p>\n    <p>Firstly, to get it out of the way, if the best design was already here, we wouldn’t be talking about how to deliver economic survivability to the world. The right solution would surely already be steamrolling over Wall Street and the Fed and the ECB and the rest.</p>`,
			Slug:    "what",
		},
		post.Post{
			ID:      "vision",
			Title:   "Vision",
			Content: "<p>In the 90s began a boom with the beginning of the commercialisation of the internet. Previous to this, it was a field primarily confined to academic institutions, where HTML, HTTP and the Web Browser technologies arose.</p><p>It took about 20 years for the race to master these technologies reached its climax with the establishment of a number of key controllers of the flow of information, Facebook, Google, Twitter, Amazon, Ebay and Paypal, among others.</p><p>The old technology is defined by its centralisation. Decentralisation has a very small part to play in it, with the use of technologies like OAUTH allowing one silo to authenticate a user to access another silo.</p><p>The databases are separated, they are curated by private interests, and dangerously, through subterfuge, these dominant silo operators have very distinct political leanings and philosophies that are not popular and are pushing an increasing number of people into a position where they have no secure accessible means of communicating information with other users, and whitewashing public opinion, or maybe it should be called ‘blue washing’ or ‘red washing’ since both colors are associated with the political groups sharing these political positions.</p>",
			Slug:    "vision",
		},
		post.Post{
			ID:      "apps",
			Title:   "Apps",
			Content: "<p>Apps text  text  text  text  text  text  text  text  text  text  text  text  text  text .</p>",
			Slug:    "apps",
		},
		post.Post{
			ID:      "development",
			Title:   "Development",
			Content: "Development Development Development Development Development Development ",
			Slug:    "development",
		},
		post.Post{
			ID:      "download",
			Title:   "Download ParallelCoin Apps",
			Content: "download download download",
			Slug:    "download",
		},
		post.Post{
			ID:      "story",
			Title:   "The true story",
			Content: "The true story The true story The true story The true storyThe true story The true story The true story The true storyThe true story The true story The true story The true storyThe true story The true story The true story The true storyThe true story The true story The true story The true story",
			Slug:    "story",
		},
		post.Post{
			ID:      "status",
			Title:   "Status",
			Content: "Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status ",
			Slug:    "status",
		},
		post.Post{
			ID:      "technology",
			Title:   "Technology",
			Content: "Technology Technology Technology Technology Technology ",
			Slug:    "technology",
		},
		post.Post{
			ID:      "documentation",
			Title:   "Documentation",
			Content: "doc doc doc doc doc doc doc doc doc ",
			Slug:    "documentation",
		},
		post.Post{
			ID:      "where",
			Title:   "Where is Parallel?",
			Content: "Where is Parallel?",
			Slug:    "where",
		},
		post.Post{
			ID:      "why",
			Title:   "Why?",
			Content: "Why?",
			Slug:    "why",
		},
	}
)
