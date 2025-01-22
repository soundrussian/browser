package devices

import "github.com/soundrussian/browser/v2/utils"

type XboxOne struct {
	p Parser
}

var (
	xboxOneName               = "Xbox One"
	xboxOneMatchRegex         = []string{`(?i)Xbox One`}
	xboxOneMatchRegexCompiled = utils.CompileRegexps(xboxOneMatchRegex)
)

func NewXboxOne(p Parser) *XboxOne {
	return &XboxOne{
		p: p,
	}
}

func (x *XboxOne) Name() string {
	return xboxOneName
}

func (x *XboxOne) Match() bool {
	return x.p.Match(xboxOneMatchRegexCompiled)
}
