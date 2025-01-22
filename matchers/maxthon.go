package matchers

import "github.com/soundrussian/browser/v2/utils"

type Maxthon struct {
	p Parser
}

var (
	maxthonName                  = "Maxthon"
	maxthonVersionRegexp         = []string{`(?i)Maxthon/([\d.]+)`}
	maxthonMatchRegex            = []string{`(i?)Maxthon`}
	maxthonVersionRegexpCompiled = utils.CompileRegexps(maxthonVersionRegexp)
	maxthonMatchRegexCompiled    = utils.CompileRegexps(maxthonMatchRegex)
)

func NewMaxthon(p Parser) *Maxthon {
	return &Maxthon{
		p: p,
	}
}

func (m *Maxthon) Name() string {
	return maxthonName
}

func (m *Maxthon) Version() string {
	return m.p.Version(maxthonVersionRegexpCompiled, 1)
}

func (m *Maxthon) Match() bool {
	return m.p.Match(maxthonMatchRegexCompiled)
}
