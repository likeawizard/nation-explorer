package gui

import (
	"likeawizard/nation-explorer/config"
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func Settings(w fyne.Window) fyne.CanvasObject {
	oldConfig, _ := config.LoadConfig()
	langOptions := config.AvailableLanguages()
	c := config.Config{
		Lang:    oldConfig.Lang,
		BaseDir: oldConfig.BaseDir,
	}

	filesOk := config.VerifyGameFiles()
	filesOkBinding := binding.BindBool(&filesOk)

	langSelect := widget.NewSelect(langOptions, func(s string) { c.Lang = s })
	langSelect.Selected = oldConfig.Lang

	dirSelect := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}
		c.BaseDir = lu.Path()
		config.SetConfig(c)
		filesOkBinding.Set(config.VerifyGameFiles())
	}, w)
	dir, err := storage.ListerForURI(storage.NewFileURI(oldConfig.BaseDir))
	if err == nil {
		dirSelect.SetLocation(dir)
	}

	fileStatus := widget.NewLabelWithData(binding.BoolToString(filesOkBinding))
	fileStatus.TextStyle = fyne.TextStyle{Bold: true}

	browseFiles := widget.NewButton("Browse", func() { dirSelect.Show() })

	settingsForm := widget.NewForm(
		widget.NewFormItem("Language", langSelect),
		widget.NewFormItem("EU4 base directory", browseFiles),
		widget.NewFormItem("Game file check:", fileStatus),
	)
	settingsForm.OnCancel = func() { w.Close() }
	settingsForm.OnSubmit = func() {

		config.SetConfig(c)
		ok := config.VerifyGameFiles()
		filesOkBinding.Set(ok)
		if ok {
			config.WriteConfig(c)
			script.LoadNationalIdeas()
			lang.ReadLangFiles()
			w.Close()
		}
		// revert to old config
		config.SetConfig(oldConfig)

	}

	settingsForm.SubmitText = "Apply"
	settingsForm.CancelText = "Close"

	return container.NewBorder(nil, nil, nil, nil, settingsForm)
}
