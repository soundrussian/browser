package platforms

import "github.com/soundrussian/browser/v2/utils"

type ChromeOS struct {
	p Parser
}

var (
	chromeOSName                  = "Chrome OS"
	chromeOSVersionRegexp         = []string{`CrOS [^\s]+ ([\d.]+)`}
	chromeOSMatchRegexp           = []string{`CrOS`}
	chromeOSVersionRegexpCompiled = utils.CompileRegexps(chromeOSVersionRegexp)
	chromeOSMatchRegexpCompiled   = utils.CompileRegexps(chromeOSMatchRegexp)
)

func NewChromeOS(p Parser) *ChromeOS {
	return &ChromeOS{
		p: p,
	}
}

func (c *ChromeOS) Name() string {
	return chromeOSName
}

func (c *ChromeOS) Version() string {
	return c.p.Version(chromeOSVersionRegexpCompiled, 1, "")
}

func (c *ChromeOS) Match() bool {
	return c.p.Match(chromeOSMatchRegexpCompiled)
}
