package matchers

import "github.com/soundrussian/browser/v2/utils"

type SogouBrowser struct {
	p Parser
}

var (
	sogouBrowserName                  = "Sogou Browser"
	sogouBrowserVersionRegexp         = []string{`(?i)(?:SogouMobileBrowser)/([\d.]+)`}
	sogouBrowserMatchRegexp           = []string{`(?i)SogouMobileBrowser`, `\bSE\b`}
	sogouBrowserVersionRegexpCompiled = utils.CompileRegexps(sogouBrowserVersionRegexp)
	sogouBrowserMatchRegexpCompiled   = utils.CompileRegexps(sogouBrowserMatchRegexp)
)

func NewSogouBrowser(p Parser) *SogouBrowser {
	return &SogouBrowser{
		p: p,
	}
}

func (s *SogouBrowser) Name() string {
	return sogouBrowserName
}

func (s *SogouBrowser) Version() string {
	return s.p.Version(sogouBrowserVersionRegexpCompiled, 1)
}

func (s *SogouBrowser) Match() bool {
	return s.p.Match(sogouBrowserMatchRegexpCompiled)
}
