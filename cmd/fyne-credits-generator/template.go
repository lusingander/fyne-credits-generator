package main

import (
	"fmt"

	credit "github.com/lusingander/fyne-credits-generator"
)

func createCreditsGo(credits []*credit.Credit) string {
	vars := ""
	for _, c := range credits {
		vars += fmt.Sprintf(`	{
		"%s",
		"%s",
		`+"`%s`"+`,
	},
`, c.Name, c.URL, c.Text)
	}
	return fmt.Sprintf(template, vars)
}

var template = `package main

import (
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func CreditsWindow(app fyne.App) fyne.Window {
	w := app.NewWindow("CREDITS")
	w.Resize(fyne.NewSize(800, 400))
	w.SetContent(CreditsContainer())
	return w
}

func CreditsContainer() fyne.CanvasObject {
	list := widget.NewVBox()
	nameLabel := widget.NewLabel("")
	urlLabel := widget.NewHyperlink("", nil)
	header := widget.NewVBox(nameLabel, urlLabel)
	entry := widget.NewMultiLineEntry()
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
	text := widget.NewScrollContainer(entry)
	license := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(header, nil, nil, nil),
		header, text,
	)
	return fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, list, nil),
		list, license,
	)
}

type credit struct {
	name, url, text string
}

var credits = []*credit{
%s}`
