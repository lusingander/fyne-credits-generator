package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	w := app.New().NewWindow("fyne-credits-generator sample")
	label := widget.NewLabel("These are the licenses for the libraries used in this repository.")
	button := widget.NewButton("Show credits", func() {
		CreditsWindow(fyne.CurrentApp(), fyne.NewSize(800, 400)).Show()
	})
	content := container.NewVBox(
		label,
		container.NewHBox(
			layout.NewSpacer(),
			button,
			layout.NewSpacer(),
		),
	)
	w.SetContent(content)
	w.ShowAndRun()
}
