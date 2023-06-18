package main

import (
	"fmt"
	"likeawizard/nation-explorer/config"
	"likeawizard/nation-explorer/gui"
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Nation Explorer")
	myWindow.Resize(fyne.NewSize(480, 480))

	mainMenu := makeMenu(myApp, myWindow)

	myWindow.SetMainMenu(mainMenu)
	myWindow.SetMaster()

	bytes, _ := os.ReadFile(config.NationalIdeaPath)

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

	myWindow.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {

	settings := func() {
		w := a.NewWindow("Settings")
		w.SetContent(gui.Settings(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	newPlayerItem := fyne.NewMenuItem("Settings", settings)
	newMenu := fyne.NewMenu("File", newPlayerItem)
	main := fyne.NewMainMenu(newMenu)
	return main
}
