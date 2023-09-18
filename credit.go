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

// Credit represents the license text, repository name and URL.
type Credit struct {
	Name, URL, Text string
}

// [`] => [` + "`" + `]
const replacedBackquote = "`" + ` + "` + "`" + `" + ` + "`"

// FormattedText returns text that replaces problematic text as code.
func (c *Credit) FormattedText() string {
	return strings.Replace(c.Text, "`", replacedBackquote, -1)
}

type collectOptions struct {
	strict bool
}

type collectOption func(*collectOptions)

func Strict(b bool) collectOption {
	return func(o *collectOptions) { o.strict = b }
}

// Collect returns the license information collected and converted to Credit type.
func Collect(options ...collectOption) ([]*Credit, error) {
	opts := &collectOptions{}
	for _, opt := range options {
		opt(opts)
	}

	buf, err := runGoCredits(opts)
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
	credits = append(credits, ownCredit)
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

func runGoCredits(opts *collectOptions) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	args := buildGoCreditsArgs(opts)
	err := gocredits.Run(args, buf, os.Stderr)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func buildGoCreditsArgs(opts *collectOptions) []string {
	args := make([]string, 0)
	if !opts.strict {
		args = append(args, "-skip-missing")
	}
	args = append(args, ".") // from current directory
	return args
}
