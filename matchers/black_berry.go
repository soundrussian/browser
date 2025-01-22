package matchers

import "github.com/soundrussian/browser/v2/utils"

type BlackBerry struct {
	p Parser
}

var (
	blackBerryName                  = "BlackBerry"
	blackBerryVersionRegexp         = []string{`BlackBerry[\da-z]+/([\d.]+)`, `Version/([\d.]+)`}
	blackBerryMatchRegex            = []string{`BlackBerry|(?i)BB10`}
	blackBerryVersionRegexpCompiled = utils.CompileRegexps(blackBerryVersionRegexp)
	blackBerryMatchRegexCompiled    = utils.CompileRegexps(blackBerryMatchRegex)
)

func NewBlackBerry(p Parser) *BlackBerry {
	return &BlackBerry{
		p: p,
	}
}

func (b *BlackBerry) Name() string {
	return blackBerryName
}

func (b *BlackBerry) Version() string {
	return b.p.Version(blackBerryVersionRegexpCompiled, 1)
}

func (b *BlackBerry) Match() bool {
	return b.p.Match(blackBerryMatchRegexCompiled)
}
