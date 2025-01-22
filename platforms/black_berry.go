package platforms

import "github.com/soundrussian/browser/v2/utils"

type BlackBerry struct {
	p Parser
}

var (
	blackBerryName                  = "BlackBerry"
	blackBerryVersionRegexp         = []string{`(?:Version|BlackBerry[\da-z]+)/([\d.]+)`}
	blackBerryMatchRegexp           = []string{`BB10|BlackBerry`}
	blackBerryVersionRegexpCompiled = utils.CompileRegexps(blackBerryVersionRegexp)
	blackBerryMatchRegexpCompiled   = utils.CompileRegexps(blackBerryMatchRegexp)
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
	return b.p.Version(blackBerryVersionRegexpCompiled, 1, "")
}

func (b *BlackBerry) Match() bool {
	return b.p.Match(blackBerryMatchRegexpCompiled)
}
