package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	credit "github.com/lusingander/fyne-credits-generator"
)

var (
	version *bool
	help    *bool
)

func parseFlags() {
	version = flag.Bool("version", false, "print version")
	help = flag.Bool("help", false, "print help")
	flag.Parse()
}

func printHelp() {
	fmt.Fprintln(os.Stderr, "Usage: fyne-credits-generator [flags]")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "fyne-credits-generator version %s\n", credit.Version)
}

func run() error {
	if *help {
		printHelp()
		return nil
	}
	if *version {
		printVersion()
		return nil
	}

	credits, err := credit.Collect()
	if err != nil {
		return err
	}
	fmt.Println(createCreditsGo(credits))
	return nil
}

func main() {
	parseFlags()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
