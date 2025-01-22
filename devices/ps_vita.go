package devices

import "github.com/soundrussian/browser/v2/utils"

type PSVita struct {
	p Parser
}

var (
	psVitaName               = "PlayStation Vita"
	psVitaMatchRegex         = []string{`(?i)PlayStation Vita`}
	psVitaMatchRegexCompiled = utils.CompileRegexps(psVitaMatchRegex)
)

func NewPSVita(p Parser) *PSVita {
	return &PSVita{
		p: p,
	}
}

func (p *PSVita) Name() string {
	return psVitaName
}

func (p *PSVita) Match() bool {
	return p.p.Match(psVitaMatchRegexCompiled)
}
