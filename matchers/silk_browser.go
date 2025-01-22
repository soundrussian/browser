package matchers

import "github.com/soundrussian/browser/v2/utils"

type SilkBrowser struct {
	p Parser
}

var (
	silkBrowserName         = "Silk"
	silkBrowserVersionRegex = []string{
		`(?i)Silk/([\d.]+)`,
		`Chrome/([\d.]+)`,
		`Safari/([\d.]+)`,
	}
	silkBrowserMatchRegex           = []string{`Silk/`}
	silkBrowserVersionRegexCompiled = utils.CompileRegexps(silkBrowserVersionRegex)
	silkBrowserMatchRegexCompiled   = utils.CompileRegexps(silkBrowserMatchRegex)
)

func NewSilkBrowser(p Parser) *SilkBrowser {
	return &SilkBrowser{
		p: p,
	}
}

func (s *SilkBrowser) Name() string {
	return silkBrowserName
}

func (s *SilkBrowser) Version() string {
	return s.p.Version(silkBrowserVersionRegexCompiled, 1)
}

func (s *SilkBrowser) Match() bool {
	return s.p.Match(silkBrowserMatchRegexCompiled)
}
