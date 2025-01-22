package matchers

import "github.com/soundrussian/browser/v2/utils"

type GoogleSearchApp struct {
	p Parser
}

var (
	googleSearchAppName = "Google Search App"
	// TODO: review this regex, prev `GSA/([\d.]+)`
	googleSearchAppVersionRegexp         = []string{`GSA/([\d]+(?:\.[\d]+)*)`}
	googleSearchAppMatchRegex            = []string{`(?i)GSA`}
	googleSearchAppVersionRegexpCompiled = utils.CompileRegexps(googleSearchAppVersionRegexp)
	googleSearchAppMatchRegexCompiled    = utils.CompileRegexps(googleSearchAppMatchRegex)
)

func NewGoogleSearchApp(p Parser) *GoogleSearchApp {
	return &GoogleSearchApp{
		p: p,
	}
}

func (g *GoogleSearchApp) Name() string {
	return googleSearchAppName
}

func (g *GoogleSearchApp) Version() string {
	return g.p.Version(googleSearchAppVersionRegexpCompiled, 1)
}

func (g *GoogleSearchApp) Match() bool {
	return g.p.Match(googleSearchAppMatchRegexCompiled)
}
