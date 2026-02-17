package matchers

import "github.com/soundrussian/browser/v2/utils"

type AlohaBrowser struct {
	p Parser
}

var (
	alohaBrowserName                 = "Aloha Browser"
	alohaBrowserVersionRegex         = []string{`(?i)AlohaBrowser/([\d.]+)`}
	alohaBrowserMatchRegex           = []string{`(?i)AlohaBrowser`}
	alohaBrowserVersionRegexCompiled = utils.CompileRegexps(alohaBrowserVersionRegex)
	alohaBrowserMatchRegexCompiled   = utils.CompileRegexps(alohaBrowserMatchRegex)
)

func NewAlohaBrowser(p Parser) *AlohaBrowser {
	return &AlohaBrowser{
		p: p,
	}
}

func (s *AlohaBrowser) Name() string {
	return alohaBrowserName
}

func (s *AlohaBrowser) Version() string {
	return s.p.Version(alohaBrowserVersionRegexCompiled, 1)
}

func (s *AlohaBrowser) Match() bool {
	return s.p.Match(alohaBrowserMatchRegexCompiled)
}
