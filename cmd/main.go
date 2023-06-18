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
	lang.ReadLangFile("data/modifiers.yml")

	fmt.Println(len(ideas))

	modifiers := make(map[string]string)

	parseGroup := func(mg script.ModGroup) {
		for _, mod := range mg.Mods {
			modifiers[mod.Name] = lang.Get(mod.Name)
		}
	}

	for _, natIdeas := range ideas {
		parseGroup(natIdeas.Ambition)
		parseGroup(natIdeas.Traditions)
		for _, idea := range natIdeas.Ideas {
			parseGroup(idea)
		}
	}

	for k, v := range modifiers {
		fmt.Println(k, v)
	}
}
