package platforms

import "github.com/soundrussian/browser/v2/utils"

type AdobeAir struct {
	p Parser
}

var (
	adobeAirName                  = "Adobe AIR"
	adobeAirVersionRegexp         = []string{`AdobeAIR/([\d.]+)`}
	adobeAirMatchRegexp           = []string{`AdobeAIR`}
	adobeAirVersionRegexpCompiled = utils.CompileRegexps(adobeAirVersionRegexp)
	adobeAirMatchRegexpCompiled   = utils.CompileRegexps(adobeAirMatchRegexp)
)

func NewAdobeAir(p Parser) *AdobeAir {
	return &AdobeAir{
		p: p,
	}
}

func (a *AdobeAir) Name() string {
	return adobeAirName
}

func (a *AdobeAir) Version() string {
	return a.p.Version(adobeAirVersionRegexpCompiled, 1, "")
}

func (a *AdobeAir) Match() bool {
	return a.p.Match(adobeAirMatchRegexpCompiled)
}
