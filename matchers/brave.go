package matchers

import "github.com/soundrussian/browser/v2/utils"

var (
	braveName                  = "Brave"
	braveVersionRegexp         = []string{`brave/([\d.]+)`}
	braveMatchRegex            = []string{`(?i)Brave`}
	braveVersionRegexpCompiled = utils.CompileRegexps(braveVersionRegexp)
	braveMatchRegexCompiled    = utils.CompileRegexps(braveMatchRegex)
)

type Brave struct {
	p Parser
}

func NewBrave(p Parser) *Brave {
	return &Brave{
		p: p,
	}
}

func (b *Brave) Name() string {
	return braveName
}

func (b *Brave) Version() string {
	return b.p.Version(braveVersionRegexpCompiled, 1)
}

func (b *Brave) Match() bool {
	return b.p.Match(braveMatchRegexCompiled)
}
