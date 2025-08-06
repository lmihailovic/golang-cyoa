package parse

import (
	"encoding/json"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func LoadStory(file string) (Story, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
		return nil, err
	}

	var story Story

	err = json.Unmarshal(bytes, &story)

	return story, nil
}
