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
	target  *string
	strict  *bool
	version *bool
	help    *bool
)

var (
	v1_3 = "1.3"
	v1_4 = "1.4"
	v2_0 = "2.0"
)

func parseFlags() {
	pname = flag.String("package", "main", "set package name")
	target = flag.String("target", v2_0, "target Fyne version (1.3|1.4|2.0)")
	strict = flag.Bool("strict", false, "error if license not found")
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

func validTarget(t string) bool {
	return t == v1_3 || t == v1_4 || t == v2_0
}

func run() error {
	if *help || !validTarget(*target) {
		printHelp()
		return nil
	}
	if *version {
		printVersion()
		return nil
	}

	credits, err := credit.Collect(credit.Strict(*strict))
	if err != nil {
		return err
	}
	fmt.Println(createCreditsGo(credits, *pname, *target))
	return nil
}

func main() {
	parseFlags()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
