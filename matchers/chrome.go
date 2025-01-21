package matchers

import "github.com/soundrussian/browser/v2/utils"

type Chrome struct {
	p Parser
}

var (
	chromeName                  = "Chrome"
	chromeVersionRegexp         = []string{`Chrome/([\d.]+)`, `CriOS/([\d.]+)`, `Safari/([\d.]+)`, `AppleWebKit/([\d.]+)`}
	chromeMatchRegex            = []string{`Chrome|CriOS`}
	chromeVersionRegexpCompiled = utils.CompileRegexps(chromeVersionRegexp)
	chromeMatchRegexCompiled    = utils.CompileRegexps(chromeMatchRegex)
)

func NewChrome(p Parser) *Chrome {
	return &Chrome{
		p: p,
	}
}

func (c *Chrome) Name() string {
	return chromeName
}

func (c *Chrome) Version() string {
	return c.p.Version(chromeVersionRegexpCompiled, 1)
}

func (c *Chrome) Match() bool {
	return c.p.Match(chromeMatchRegexCompiled)
}
