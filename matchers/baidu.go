package matchers

import "github.com/soundrussian/browser/v2/utils"

type Baidu struct {
	p Parser
}

var (
	baiduName                  = "Baidu"
	baiduVersionRegexp         = []string{`baiduboxapp/([\d.]+)`}
	baiduMatchRegex            = []string{`(?i)baiduboxapp`}
	baiduVersionRegexpCompiled = utils.CompileRegexps(baiduVersionRegexp)
	baiduMatchRegexCompiled    = utils.CompileRegexps(baiduMatchRegex)
)

func NewBaidu(p Parser) *Baidu {
	return &Baidu{
		p: p,
	}
}

func (a *Baidu) Name() string {
	return baiduName
}

func (a *Baidu) Version() string {
	return a.p.Version(baiduVersionRegexpCompiled, 1)
}

func (a *Baidu) Match() bool {
	return a.p.Match(baiduMatchRegexCompiled)
}
