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

	//for _, v := range story {
	//	fmt.Printf("\n\n== %v ==\n", v.Title)
	//	for _, s := range v.Story {
	//		println(s)
	//	}
	//	for index, o := range v.Options {
	//		fmt.Printf("\nPress %v to %v", index+1, o.Text)
	//	}
	//}

	return story, nil
}
