package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	credit "github.com/lusingander/fyne-credits-generator"
)

var (
	pname   *string
	old     *bool
	version *bool
	help    *bool
)

func parseFlags() {
	pname = flag.String("package", "main", "set package name")
	old = flag.Bool("old", false, "old style (Fyne v1.3.x or earlier)")
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
	fmt.Println(createCreditsGo(credits, *pname, *old))
	return nil
}

func main() {
	parseFlags()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
