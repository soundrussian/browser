package devices

import "github.com/soundrussian/browser/v2/utils"

type Iphone struct {
	p Parser
}

var (
	iPhoneName               = "iPhone"
	iPhoneMatchRegex         = []string{`iPhone`}
	iPhoneMatchRegexCompiled = utils.CompileRegexps(iPhoneMatchRegex)
)

func NewIphone(p Parser) *Iphone {
	return &Iphone{
		p: p,
	}
}

func (i *Iphone) Name() string {
	return iPhoneName
}

func (i *Iphone) Match() bool {
	return i.p.Match(iPhoneMatchRegexCompiled)
}
