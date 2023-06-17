package main

import (
	"fmt"
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"
	"os"
)

const (
	GAME_PATH = "/home/arturs/.local/share/Steam/steamapps/common/Europa Universalis IV/common/ideas/00_country_ideas.txt"
	IDEA_PATH = "common/ideas/00_country_ideas.txt"
)

func main() {
	bytes, err := os.ReadFile(GAME_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	ideas := script.ParseIdeas(bytes)

	fmt.Println(len(ideas))

	for k, _ := range ideas {
		fmt.Println(lang.Get(k))
	}

}
