package main

import (
	"fmt"
	"likeawizard/nation-explorer/config"
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"
	"os"
)

func main() {
	bytes, err := os.ReadFile(config.NationalIdeaPath)
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
