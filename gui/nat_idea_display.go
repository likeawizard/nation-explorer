package gui

import (
	"fmt"
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NationalIdeaDisplay struct {
	title     *widget.Label
	tradition updateFn
	ambition  updateFn
	ideas     []updateFn
	obj       fyne.CanvasObject
}

type updateFn func(script.ModGroup)

func modGroupDisplay() (fyne.CanvasObject, updateFn) {
	label := widget.NewLabel("")
	label.TextStyle = fyne.TextStyle{Bold: true}
	mods := widget.NewLabel("")
	c := container.NewVBox(label, mods)
	f := func(mg script.ModGroup) {
		label.SetText(lang.Get(mg.Name))
		mods.SetText(modGroupToString(mg))
	}

	return c, f
}

func NewNationalIDeaDisplay() *NationalIdeaDisplay {
	title := widget.NewLabel("")
	title.TextStyle = fyne.TextStyle{Bold: true}

	tc, tf := modGroupDisplay()
	ac, af := modGroupDisplay()
	ics := make([]fyne.CanvasObject, 7)
	ifs := make([]updateFn, 7)
	for i := 0; i < 7; i++ {
		ideaComp, ideaFn := modGroupDisplay()
		ics[i] = ideaComp
		ifs[i] = ideaFn
	}
	tradAndIdeas := container.NewHBox(tc, ac)
	il := container.NewVBox(ics...)
	border := container.NewBorder(title, nil, nil, container.NewVBox(tradAndIdeas, il))
	center := container.NewCenter(border)
	return &NationalIdeaDisplay{
		title:     title,
		tradition: tf,
		ambition:  af,
		ideas:     ifs,

		obj: center}
}

func (nid *NationalIdeaDisplay) Set(id string) {
	ni := script.GetNatIdeas()[id]
	nid.title.SetText(lang.Get(ni.Name))
	nid.tradition(ni.Traditions)
	nid.ambition(ni.Ambition)
	for i := range ni.Ideas {
		nid.ideas[i](ni.Ideas[i])
	}
}

func modGroupToString(mg script.ModGroup) string {
	s := ""
	for _, v := range mg.Mods {
		s += fmt.Sprintf("%s %s\n", v.Value, lang.Get(v.Name))
	}
	return s
}
