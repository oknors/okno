package mod

type Cache struct {
	Response bool        `json:"r"`
	Data     interface{} `json:"d"`
}
type WebSite struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
