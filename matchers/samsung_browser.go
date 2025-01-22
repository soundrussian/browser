package matchers

import "github.com/soundrussian/browser/v2/utils"

type SamsungBrowser struct {
	p Parser
}

var (
	samsungBrowserName         = "Samsung Browser"
	samsungBrowserVersionRegex = []string{
		`(?i)SamsungBrowser/([\d.]+)`,
		`Chrome/([\d.]+)`,
		`CriOS/([\d.]+)`,
		`Safari/([\d.]+)`,
		`AppleWebKit/([\d.]+)`,
	}
	samsungBrowserMatchRegex           = []string{`SamsungBrowser`}
	samsungBrowserVersionRegexCompiled = utils.CompileRegexps(samsungBrowserVersionRegex)
	samsungBrowserMatchRegexCompiled   = utils.CompileRegexps(samsungBrowserMatchRegex)
)

func NewSamsungBrowser(p Parser) *SamsungBrowser {
	return &SamsungBrowser{
		p: p,
	}
}

func (s *SamsungBrowser) Name() string {
	return samsungBrowserName
}

func (s *SamsungBrowser) Version() string {
	return s.p.Version(samsungBrowserVersionRegexCompiled, 1)
}

func (s *SamsungBrowser) Match() bool {
	return s.p.Match(samsungBrowserMatchRegexCompiled)
}
