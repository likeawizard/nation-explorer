package gui

import (
	"likeawizard/nation-explorer/lang"
	"likeawizard/nation-explorer/script"
	"math/rand"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

func ListNationalIdeas() fyne.CanvasObject {
	ideaDisplay := NewNationalIDeaDisplay()

	natIdeaIds := []string{}
	for k := range script.GetNatIdeas() {
		natIdeaIds = append(natIdeaIds, k)
	}
	sortByLabel(natIdeaIds)
	ideaBinding := binding.BindStringList(&natIdeaIds)

	search := widget.NewEntry()
	search.PlaceHolder = "Search..."
	search.OnChanged = bindSearch(ideaBinding)

	list := widget.NewListWithData(ideaBinding,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			s, _ := i.(binding.String).Get()
			o.(*widget.Label).SetText(lang.Get(s))
		})
	list.OnSelected = func(id widget.ListItemID) {
		i, _ := ideaBinding.GetItem(id)
		s, _ := i.(binding.String).Get()
		ideaDisplay.Set(s)
	}

	selectRandom := widget.NewButton("Random", func() {
		list.Select(rand.Intn(list.Length()))

	})

	ideaList := container.NewBorder(search, selectRandom, nil, nil, list)

	return container.NewHSplit(ideaList, container.NewVScroll(ideaDisplay.obj))
}

func bindSearch(data binding.ExternalStringList) func(string) {
	return func(s string) {
		filteredIdeas := []string{}
		for k := range script.GetNatIdeas() {
			if strings.Contains(removeSpecialChars(strings.ToLower(lang.Get(k))), removeSpecialChars(strings.ToLower(s))) {
				filteredIdeas = append(filteredIdeas, k)
			}
		}

		sortByLabel(filteredIdeas)
		data.Set(filteredIdeas)
	}
}

// Remove accented and other language specific chars from strings
func removeSpecialChars(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	r, _, _ := transform.String(t, s)
	return r
}

func sortByLabel(s []string) {
	sort.Slice(s, func(i, j int) bool {
		return lang.Get(s[i]) < lang.Get(s[j])
	})
}
