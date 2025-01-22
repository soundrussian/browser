package devices

import "github.com/soundrussian/browser/v2/utils"

type Switch struct {
	p Parser
}

var (
	switchName               = "Nintendo Switch"
	switchMatchRegex         = []string{`(?i)Nintendo Switch`}
	switchMatchRegexCompiled = utils.CompileRegexps(switchMatchRegex)
)

func NewSwitch(p Parser) *Switch {
	return &Switch{
		p: p,
	}
}

func (s *Switch) Name() string {
	return switchName
}

func (s *Switch) Match() bool {
	return s.p.Match(switchMatchRegexCompiled)
}
