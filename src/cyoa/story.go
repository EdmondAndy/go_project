package cyoa

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Demo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
