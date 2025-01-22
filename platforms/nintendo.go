package platforms

import "github.com/soundrussian/browser/v2/utils"

type Nintendo struct {
	p Parser
}

var (
	nintendoName               = "Nintendo"
	nintendoMatchRegex         = []string{`Nintendo`}
	nintendoMatchRegexCompiled = utils.CompileRegexps(nintendoMatchRegex)
)

func NewNintendo(p Parser) *Nintendo {
	return &Nintendo{
		p: p,
	}
}

func (k *Nintendo) Name() string {
	return nintendoName
}

// Version returns empty string for Nintendo
func (k *Nintendo) Version() string {
	return ""
}

func (k *Nintendo) Match() bool {
	return k.p.Match(nintendoMatchRegexCompiled)
}
