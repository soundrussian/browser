package matchers

import "github.com/soundrussian/browser/v2/utils"

type Otter struct {
	p Parser
}

var (
	otterName                  = "Otter"
	otterVersionRegexp         = []string{`Otter/([\d.]+)`}
	otterMatchRegexp           = []string{`Otter`}
	otterVersionRegexpCompiled = utils.CompileRegexps(otterVersionRegexp)
	otterMatchRegexpCompiled   = utils.CompileRegexps(otterMatchRegexp)
)

func NewOtter(p Parser) *Otter {
	return &Otter{
		p: p,
	}
}

func (o *Otter) Name() string {
	return otterName
}

func (o *Otter) Version() string {
	return o.p.Version(otterVersionRegexpCompiled, 1)
}

func (o *Otter) Match() bool {
	return o.p.Match(otterMatchRegexpCompiled)
}
