package matchers

import "github.com/soundrussian/browser/v2/utils"

var (
	vivaldiName                  = "Vivaldi"
	vivaldiVersionRegexp         = []string{`Vivaldi/([\d.]+)`}
	vivaldiMatchRegex            = []string{`(?i)Vivaldi`}
	vivaldiVersionRegexpCompiled = utils.CompileRegexps(vivaldiVersionRegexp)
	vivaldiMatchRegexCompiled    = utils.CompileRegexps(vivaldiMatchRegex)
)

type Vivaldi struct {
	p Parser
}

func NewVivaldi(p Parser) *Vivaldi {
	return &Vivaldi{
		p: p,
	}
}

func (v *Vivaldi) Name() string {
	return vivaldiName
}

func (v *Vivaldi) Version() string {
	return v.p.Version(vivaldiVersionRegexpCompiled, 1)
}

func (v *Vivaldi) Match() bool {
	return v.p.Match(vivaldiMatchRegexCompiled)
}
