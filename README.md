fyne-credits-generator
====

Automatically generate credits window for the [Fyne](https://fyne.io/) application.

## Installation

```
$ go get github.com/lusingander/fyne-credits-generator/cmd/fyne-credits-generator
```

## Usage

First, execute `fyne-credits-generator` command in the directory where go.mod is located.

```
$ cd <fyne project directory>
$ fyne-credits-generator > credits.go
```

Then, all you have to do is call `CreditsWindow` function in your Fyne application, like this:

```go
button := widget.NewButton("Show credits", func() {
    CreditsWindow(fyne.CurrentApp(), fyne.NewSize(800, 400)).Show()
})
```
And you can show a window like this:
<img src="./resource/screenshot.png">

See the [sample application](./cmd/sample) for an example.

More info:
```
$ fyne-credits-generator -help
```

## Requirements

Application must use go modules for dependency management.

## Note

This application uses [gocredits](https://github.com/Songmu/gocredits) to collect LICENSE files.
