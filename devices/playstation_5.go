package devices

import "github.com/soundrussian/browser/v2/utils"

type PlayStation5 struct {
	p Parser
}

var (
	playStation5Name               = "PlayStation 5"
	playStation5MatchRegex         = []string{`(?i)PlayStation 5`}
	playStation5MatchRegexCompiled = utils.CompileRegexps(playStation5MatchRegex)
)

func NewPlayStation5(p Parser) *PlayStation5 {
	return &PlayStation5{
		p: p,
	}
}

func (p *PlayStation5) Name() string {
	return playStation5Name
}

func (p *PlayStation5) Match() bool {
	return p.p.Match(playStation5MatchRegexCompiled)
}
