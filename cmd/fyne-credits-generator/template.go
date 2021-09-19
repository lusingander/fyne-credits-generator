package main

import (
	"fmt"

	credit "github.com/lusingander/fyne-credits-generator"
)

func createCreditsGo(credits []*credit.Credit, pname string, target string) string {
	vars := ""
	for _, c := range credits {
		vars += fmt.Sprintf(`	{
		"%s",
		"%s",
		`+"`%s`"+`,
	},
`, c.Name, c.URL, c.FormattedText())
	}
	var importTemplate, creditsContainerTemplate string
	if target == v1_3 {
		importTemplate = v1_3StyleImportTemplate
		creditsContainerTemplate = v1_3StyleCreditsContainerTemplate
	} else if target == v1_4 {
		importTemplate = v1_4StyleImportTemplate
		creditsContainerTemplate = v1_4StyleCreditsContainerTemplate
	} else { // v2_0
		importTemplate = v2_0StyleImportTemplate
		creditsContainerTemplate = v2_0StyleCreditsContainerTemplate
	}
	return fmt.Sprintf(baseTemplate, credit.Version, pname, importTemplate, creditsContainerTemplate, vars)
}

var baseTemplate = `// Code generated by github.com/lusingander/fyne-credits-generator (v%s); DO NOT EDIT.

package %s

%s

// CreditsWindow returns a window displaying a list of licenses.
func CreditsWindow(app fyne.App, size fyne.Size) fyne.Window {
	w := app.NewWindow("CREDITS")
	w.Resize(size)
	w.SetContent(CreditsContainer())
	return w
}

%s

type credit struct {
	name, url, text string
}

var credits = []*credit{
%s}`

var v1_3StyleImportTemplate = `import (
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)`

var v1_4StyleImportTemplate = `import (
	"net/url"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)`

var v2_0StyleImportTemplate = `import (
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)`

var v1_3StyleCreditsContainerTemplate = `// CreditsContainer returns a container displaying a list of licenses.
func CreditsContainer() fyne.CanvasObject {
	list := widget.NewVBox()
	nameLabel := widget.NewLabel("")
	urlLabel := widget.NewHyperlink("", nil)
	header := widget.NewVBox(nameLabel, urlLabel)
	entry := widget.NewMultiLineEntry()
	entry.Wrapping = fyne.TextWrapBreak
	for _, c := range credits {
		c := c
		button := widget.NewButton(c.name, func() {
			nameLabel.SetText(c.name)
			u, _ := url.Parse(c.url)
			urlLabel.SetText(c.url)
			urlLabel.SetURL(u)
			entry.SetText(c.text)
		})
		list.Append(button)
	}
	listContainer := widget.NewVScrollContainer(list)
	text := widget.NewScrollContainer(entry)
	license := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(header, nil, nil, nil),
		header, text,
	)
	splitContainer := widget.NewHSplitContainer(listContainer, license)
	splitContainer.SetOffset(0)
	return splitContainer
}`

var v1_4StyleCreditsContainerTemplate = `// CreditsContainer returns a container displaying a list of licenses.
func CreditsContainer() fyne.CanvasObject {
	nameLabel := widget.NewLabel("")
	urlLabel := widget.NewHyperlink("", nil)
	header := widget.NewVBox(nameLabel, urlLabel)
	entry := widget.NewMultiLineEntry()
	entry.Wrapping = fyne.TextWrapBreak
	width := 0
	for _, c := range credits {
		l := len(c.name)
		if l > width {
			width = l
		}
	}
	list := widget.NewList(
		func() int {
			return len(credits)
		},
		func() fyne.CanvasObject {
			dummy := strings.Repeat("*", width)
			label := widget.NewLabel(dummy)
			label.Alignment = fyne.TextAlignCenter
			return label
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(credits[id].name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		c := credits[id]
		nameLabel.SetText(c.name)
		u, _ := url.Parse(c.url)
		urlLabel.SetText(c.url)
		urlLabel.SetURL(u)
		entry.SetText(c.text)
	}
	list.Select(0)
	text := widget.NewScrollContainer(entry)
	license := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(header, nil, nil, nil),
		header, text,
	)
	splitContainer := widget.NewHSplitContainer(list, license)
	splitContainer.SetOffset(0)
	return splitContainer
}`

var v2_0StyleCreditsContainerTemplate = `// CreditsContainer returns a container displaying a list of licenses.
func CreditsContainer() fyne.CanvasObject {
	nameLabel := widget.NewLabel("")
	urlLabel := widget.NewHyperlink("", nil)
	header := container.NewVBox(nameLabel, urlLabel)
	entry := widget.NewMultiLineEntry()
	entry.Wrapping = fyne.TextWrapBreak
	width := 0
	for _, c := range credits {
		l := len(c.name)
		if l > width {
			width = l
		}
	}
	list := widget.NewList(
		func() int {
			return len(credits)
		},
		func() fyne.CanvasObject {
			dummy := strings.Repeat("*", width)
			label := widget.NewLabel(dummy)
			label.Alignment = fyne.TextAlignCenter
			return label
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(credits[id].name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		c := credits[id]
		nameLabel.SetText(c.name)
		u, _ := url.Parse(c.url)
		urlLabel.SetText(c.url)
		urlLabel.SetURL(u)
		entry.SetText(c.text)
	}
	list.Select(0)
	text := container.NewScroll(entry)
	license := container.New(
		layout.NewBorderLayout(header, nil, nil, nil),
		header, text,
	)
	splitContainer := container.NewHSplit(list, license)
	splitContainer.SetOffset(0)
	return splitContainer
}`
