package matchers

import "github.com/soundrussian/browser/v2/utils"

type Firefox struct {
	p Parser
}

var (
	firefoxName                  = "Firefox"
	firefoxVersionRegexp         = []string{`Firefox/([\d.]+)`, `FxiOS/([\d.]+)`}
	firefoxMatchRegex            = []string{`Firefox`, `FxiOS`}
	firefoxVersionRegexpCompiled = utils.CompileRegexps(firefoxVersionRegexp)
	firefoxMatchRegexCompiled    = utils.CompileRegexps(firefoxMatchRegex)
)

func NewFirefox(p Parser) *Firefox {
	return &Firefox{
		p: p,
	}
}

func (f *Firefox) Name() string {
	return firefoxName
}

func (f *Firefox) Version() string {
	return f.p.Version(firefoxVersionRegexpCompiled, 1)
}

func (f *Firefox) Match() bool {
	return f.p.Match(firefoxMatchRegexCompiled)
}
