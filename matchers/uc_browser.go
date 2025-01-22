package matchers

import "github.com/soundrussian/browser/v2/utils"

type UCBrowser struct {
	p Parser
}

var (
	ucBrowserName                  = "UCBrowser"
	ucBrowserVersionRegexp         = []string{`UCBrowser/([\d.]+)`, `UCWEB(?:/)?([\d.]+)`}
	ucBrowserMatchRegexp           = []string{`UC(Browser|WEB)`}
	ucBrowserVersionRegexpCompiled = utils.CompileRegexps(ucBrowserVersionRegexp)
	ucBrowserMatchRegexpCompiled   = utils.CompileRegexps(ucBrowserMatchRegexp)
)

func NewUCBrowser(p Parser) *UCBrowser {
	return &UCBrowser{
		p: p,
	}
}

func (u *UCBrowser) Name() string {
	return ucBrowserName
}

func (u *UCBrowser) Version() string {
	return u.p.Version(ucBrowserVersionRegexpCompiled, 1)
}

func (u *UCBrowser) Match() bool {
	return u.p.Match(ucBrowserMatchRegexpCompiled)
}
