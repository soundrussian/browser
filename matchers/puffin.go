package matchers

import "github.com/soundrussian/browser/v2/utils"

type Puffin struct {
	p Parser
}

var (
	puffinName               = "Puffin"
	puffinVersionReg         = []string{`(?i)Puffin/([\d.]+)`}
	puffinMatchRegex         = []string{`(?i)Puffin`}
	puffinVersionRegCompiled = utils.CompileRegexps(puffinVersionReg)
	puffinMatchRegexCompiled = utils.CompileRegexps(puffinMatchRegex)
)

func NewPuffin(p Parser) *Puffin {
	return &Puffin{
		p: p,
	}
}

func (pu *Puffin) Name() string {
	return puffinName
}

func (pu *Puffin) Version() string {
	return pu.p.Version(puffinVersionRegCompiled, 1)
}

func (pu *Puffin) Match() bool {
	return pu.p.Match(puffinMatchRegexCompiled)
}
