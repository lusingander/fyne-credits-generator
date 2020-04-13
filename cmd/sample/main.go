package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	w := app.New().NewWindow("fyne-credits-generator sample")
	label := widget.NewLabel("These are the licenses for the libraries used in this repository.")
	button := widget.NewButton("Show credits", func() {
		CreditsWindow(fyne.CurrentApp(), fyne.NewSize(800, 400)).Show()
	})
	content := widget.NewVBox(
		label,
		widget.NewHBox(
			layout.NewSpacer(),
			button,
			layout.NewSpacer(),
		),
	)
	w.SetContent(content)
	w.ShowAndRun()
}
