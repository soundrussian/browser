package matchers

import "github.com/soundrussian/browser/v2/utils"

type Alipay struct {
	p Parser
}

var (
	aliPayName                  = "Alipay"
	aliPayVersionRegexp         = []string{`AlipayClient(?:HK)?/([\d.]+)`}
	aliPayMatchRegex            = []string{`(?i)Alipay`}
	aliPayVersionRegexpCompiled = utils.CompileRegexps(aliPayVersionRegexp)
	aliPayMatchRegexCompiled    = utils.CompileRegexps(aliPayMatchRegex)
)

func NewAlipay(p Parser) *Alipay {
	return &Alipay{
		p: p,
	}
}

func (a *Alipay) Name() string {
	return aliPayName
}

func (a *Alipay) Version() string {
	return a.p.Version(aliPayVersionRegexpCompiled, 1)
}

func (a *Alipay) Match() bool {
	return a.p.Match(aliPayMatchRegexCompiled)
}
