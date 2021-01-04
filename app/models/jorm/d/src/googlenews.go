package dsrc

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type GoogleNewsRss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Media   string   `xml:"media,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Language      string `xml:"language"`
		WebMaster     string `xml:"webMaster"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Description   string `xml:"description"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Description struct {
				Text string `xml:",chardata"`
				A    struct {
					Text   string `xml:",chardata"`
					Href   string `xml:"href,attr"`
					Target string `xml:"target,attr"`
				} `xml:"a"`
				Font struct {
					Text  string `xml:",chardata"`
					Color string `xml:"color,attr"`
				} `xml:"font"`
				P string `xml:"p"`
			} `xml:"description"`
			Source struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"source"`
			Content struct {
				Text   string `xml:",chardata"`
				URL    string `xml:"url,attr"`
				Medium string `xml:"medium,attr"`
				Width  string `xml:"width,attr"`
				Height string `xml:"height,attr"`
			} `xml:"content"`
		} `xml:"item"`
	} `xml:"channel"`
}

func getGoogleNewsRss() {
	// var newsRaw GoogleNewsRss
	resp, err := http.Get("https://news.google.com/rss/search?q=parallelcoin")
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	rss := GoogleNewsRss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return
	}

	fmt.Printf("Channel title: %v\n", rss.Channel.Title)
	fmt.Printf("Channel link: %v\n", rss.Channel.Link)

	for i, item := range rss.Channel.Item {
		fmt.Printf("%v. item title: %v\n", i, item.Title)
	}
	fmt.Println("sssssssssssssssssssssasas", resp)
}
