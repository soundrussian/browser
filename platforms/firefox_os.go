package platforms

import "github.com/soundrussian/browser/v2/utils"

type FirefoxOS struct {
	p Parser
}

var (
	firefoxOSName                      = "Firefox OS"
	firefoxOSMatchRegexp               = []string{`Firefox`}
	firefoxDeviceExcludeRegexp         = []string{`Android|Linux|BlackBerry|Windows|Mac`}
	firefoxOSMatchRegexpCompiled       = utils.CompileRegexps(firefoxOSMatchRegexp)
	firefoxDeviceExcludeRegexpCompiled = utils.CompileRegexps(firefoxDeviceExcludeRegexp)
)

func NewFirefoxOS(p Parser) *FirefoxOS {
	return &FirefoxOS{
		p: p,
	}
}

func (f *FirefoxOS) Name() string {
	return firefoxOSName
}

func (f *FirefoxOS) Version() string {
	return "0"
}

func (f *FirefoxOS) Match() bool {
	return !f.isExcludeDevice() && f.p.Match(firefoxOSMatchRegexpCompiled)
}

func (f *FirefoxOS) isExcludeDevice() bool {
	return f.p.Match(firefoxDeviceExcludeRegexpCompiled)
}
