package mod

type News struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Date  string `json:"date"`
}

type VNews struct {
	News []News `json:"news"`
}
