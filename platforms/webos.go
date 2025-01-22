package platforms

import "github.com/soundrussian/browser/v2/utils"

type WebOS struct {
	p Parser
}

var (
	webOSName                  = "WebOS"
	webOSVersionRegexp         = []string{`OS\/?([\d.]+)`}
	webOSRegex                 = []string{`(?i)WebOS|hpwOS`}
	webOSVersionRegexpCompiled = utils.CompileRegexps(webOSVersionRegexp)
	webOSRegexCompiled         = utils.CompileRegexps(webOSRegex)
)

func NewWebOS(p Parser) *WebOS {
	return &WebOS{
		p: p,
	}
}

func (k *WebOS) Name() string {
	return webOSName
}

// Version returns version for WebOS
func (k *WebOS) Version() string {
	return k.p.Version(webOSVersionRegexpCompiled, 1, "")
}

func (k *WebOS) Match() bool {
	return k.p.Match(webOSRegexCompiled)
}
