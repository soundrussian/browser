package matchers

import "github.com/soundrussian/browser/v2/utils"

type DuckDuckGo struct {
	p Parser
}

var (
	duckDuckGoName                  = "DuckDuckGo"
	duckDuckGoVersionRegexp         = []string{`DuckDuck(?:Go|GoKite)?/([\d.]+)`}
	duckDuckGoMatchRegex            = []string{`DuckDuck(Go|GoKite)?`}
	duckDuckGoVersionRegexpCompiled = utils.CompileRegexps(duckDuckGoVersionRegexp)
	duckDuckGoMatchRegexCompiled    = utils.CompileRegexps(duckDuckGoMatchRegex)
)

func NewDuckDuckGo(p Parser) *DuckDuckGo {
	return &DuckDuckGo{
		p: p,
	}
}

func (ddg *DuckDuckGo) Name() string {
	return duckDuckGoName
}

func (ddg *DuckDuckGo) Version() string {
	return ddg.p.Version(duckDuckGoVersionRegexpCompiled, 1)
}

func (ddg *DuckDuckGo) Match() bool {
	return ddg.p.Match(duckDuckGoMatchRegexCompiled)
}
