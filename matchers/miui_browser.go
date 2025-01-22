package matchers

import "github.com/soundrussian/browser/v2/utils"

type MiuiBrowser struct {
	p Parser
}

var (
	miuiBrowserName                  = "Miui Browser"
	miuiBrowserVersionRegexp         = []string{`MiuiBrowser/([\d.]+)`}
	miuiBrowserMatchRegexp           = []string{`MiuiBrowser`}
	miuiBrowserVersionRegexpCompiled = utils.CompileRegexps(miuiBrowserVersionRegexp)
	miuiBrowserMatchRegexpCompiled   = utils.CompileRegexps(miuiBrowserMatchRegexp)
)

func NewMiuiBrowser(p Parser) *MiuiBrowser {
	return &MiuiBrowser{
		p: p,
	}
}

func (m *MiuiBrowser) Name() string {
	return miuiBrowserName
}

func (m *MiuiBrowser) Version() string {
	return m.p.Version(miuiBrowserVersionRegexpCompiled, 1)
}

func (m *MiuiBrowser) Match() bool {
	return m.p.Match(miuiBrowserMatchRegexpCompiled)
}
