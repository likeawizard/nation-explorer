package main

import (
	"likeawizard/nation-explorer/gui"
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Nation Explorer")
	myWindow.Resize(fyne.NewSize(1024, 800))

	mainMenu := makeMenu(myApp, myWindow)

	myWindow.SetMainMenu(mainMenu)
	myWindow.SetMaster()

	script.LoadNationalIdeas()
	lang.ReadLangFile("data/modifiers.yml")

	myWindow.SetContent(gui.ListNationalIdeas())
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
