package matchers

import "github.com/soundrussian/browser/v2/utils"

type Safari struct {
	p Parser
}

var (
	safariName                  = "Safari"
	safariVersionRegexp         = []string{`Version/([\d.]+)`, `Safari/([\d.]+)`, `AppleWebKit/([\d.]+)`}
	safariMatchRegex            = []string{`Safari`}
	safariVersionRegexpCompiled = utils.CompileRegexps(safariVersionRegexp)
	safariMatchRegexCompiled    = utils.CompileRegexps(safariMatchRegex)
)

func NewSafari(p Parser) *Safari {
	return &Safari{
		p: p,
	}
}

func (s *Safari) Name() string {
	return safariName
}

func (s *Safari) Version() string {
	return s.p.Version(safariVersionRegexpCompiled, 1)
}

func (s *Safari) Match() bool {
	return s.p.Match(safariMatchRegexCompiled)
}
