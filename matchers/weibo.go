package matchers

import "github.com/soundrussian/browser/v2/utils"

type Weibo struct {
	p Parser
}

var (
	weiboName                  = "Weibo"
	weiboVersionRegexp         = []string{`(?i)(?:__weibo__)([\d.]+)`}
	weiboMatchRegexp           = []string{`__weibo__`}
	weiboVersionRegexpCompiled = utils.CompileRegexps(weiboVersionRegexp)
	weiboMatchRegexpCompiled   = utils.CompileRegexps(weiboMatchRegexp)
)

func NewWeibo(p Parser) *Weibo {
	return &Weibo{
		p: p,
	}
}

func (w *Weibo) Name() string {
	return weiboName
}

func (w *Weibo) Version() string {
	return w.p.Version(weiboVersionRegexpCompiled, 1)
}

func (w *Weibo) Match() bool {
	return w.p.Match(weiboMatchRegexpCompiled)
}
