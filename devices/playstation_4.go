package devices

import "github.com/soundrussian/browser/v2/utils"

type PlayStation4 struct {
	p Parser
}

var (
	playStation4Name               = "PlayStation 4"
	playStation4MatchRegex         = []string{`(?i)PLAYSTATION 4`}
	playStation4MatchRegexCompiled = utils.CompileRegexps(playStation4MatchRegex)
)

func NewPlayStation4(p Parser) *PlayStation4 {
	return &PlayStation4{
		p: p,
	}
}

func (p *PlayStation4) Name() string {
	return playStation4Name
}

func (p *PlayStation4) Match() bool {
	return p.p.Match(playStation4MatchRegexCompiled)
}
