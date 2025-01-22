package matchers

import "github.com/soundrussian/browser/v2/utils"

type PaleMoon struct {
	p Parser
}

var (
	paleMoonName                  = "Pale Moon"
	paleMoonVersionRegexp         = []string{`PaleMoon/([\d.]+)`}
	paleMoonMatchRegex            = []string{`PaleMoon`}
	paleMoonVersionRegexpCompiled = utils.CompileRegexps(paleMoonVersionRegexp)
	paleMoonMatchRegexCompiled    = utils.CompileRegexps(paleMoonMatchRegex)
)

func NewPaleMoon(p Parser) *PaleMoon {
	return &PaleMoon{
		p: p,
	}
}

func (pa *PaleMoon) Name() string {
	return paleMoonName
}

func (pa *PaleMoon) Version() string {
	return pa.p.Version(paleMoonVersionRegexpCompiled, 1)
}

func (pa *PaleMoon) Match() bool {
	return pa.p.Match(paleMoonMatchRegexCompiled)
}
