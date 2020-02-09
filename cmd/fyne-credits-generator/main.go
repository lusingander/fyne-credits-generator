package main

import (
	"fmt"
	"log"

	credit "github.com/lusingander/fyne-credits-generator"
)

func run() error {
	credits, err := credit.Collect()
	if err != nil {
		return err
	}
	fmt.Println(createCreditsGo(credits))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
