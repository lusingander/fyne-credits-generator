fyne-credits-generaotor
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
    CreditsWindow(fyne.CurrentApp()).Show()
})
```
See the [sample application](./cmd/sample) for an example.

## Requirements

Application must use go modules for dependency management.

## Note

This application uses [gocredits](https://github.com/Songmu/gocredits) to collect LICENSE files.
