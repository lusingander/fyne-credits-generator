package credit

import (
	"bytes"
	"os"
	"strings"

	"github.com/Songmu/gocredits"
)

const (
	splitterLicense = "================================================================"
	splitterText    = "----------------------------------------------------------------"
)

type Credit struct {
	Name, URL, Text string
}

func Collect() ([]*Credit, error) {
	buf, err := runGoCredits()
	if err != nil {
		return nil, err
	}
	licenses := strings.Split(buf.String(), splitterLicense)
	credits := make([]*Credit, 0)
	for _, l := range licenses {
		c := newCredit(l)
		if c != nil {
			credits = append(credits, c)
		}
	}
	return credits, nil
}

func newCredit(text string) *Credit {
	l := strings.Split(text, splitterText)
	if len(l) < 2 {
		return nil
	}
	s := strings.Split(strings.Trim(l[0], "\n"), "\n")
	return &Credit{
		Name: s[0],
		URL:  s[1],
		Text: l[1],
	}
}

func runGoCredits() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	err := gocredits.Run([]string{"."} /* from current directory */, buf, os.Stderr)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
