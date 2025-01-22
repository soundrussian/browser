package devices

import "github.com/soundrussian/browser/v2/utils"

type BlackberryPlaybook struct {
	p Parser
}

var (
	bbPlaybookName               = "BlackBerry Playbook"
	bbPlaybookMatchRegex         = []string{`PlayBook.*?RIM Tablet`}
	bbPlaybookMatchRegexCompiled = utils.CompileRegexps(bbPlaybookMatchRegex)
)

func NewBlackberryPlaybook(p Parser) *BlackberryPlaybook {
	return &BlackberryPlaybook{
		p: p,
	}
}

func (b *BlackberryPlaybook) Name() string {
	return bbPlaybookName
}

func (b *BlackberryPlaybook) Match() bool {
	return b.p.Match(bbPlaybookMatchRegexCompiled)
}
