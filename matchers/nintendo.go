package matchers

import "github.com/soundrussian/browser/v2/utils"

type Nintendo struct {
	p Parser
}

var (
	nintendoName                  = "Nintendo"
	nintendoVersionRegexp         = []string{`NintendoBrowser/([\d.]+)`}
	nintendoMatchRegexp           = []string{`NintendoBrowser`}
	nintendoVersionRegexpCompiled = utils.CompileRegexps(nintendoVersionRegexp)
	nintendoMatchRegexpCompiled   = utils.CompileRegexps(nintendoMatchRegexp)
)

func NewNintendo(p Parser) *Nintendo {
	return &Nintendo{
		p: p,
	}
}

func (n *Nintendo) Name() string {
	return nintendoName
}

func (n *Nintendo) Version() string {
	return n.p.Version(nintendoVersionRegexpCompiled, 1)
}

func (n *Nintendo) Match() bool {
	return n.p.Match(nintendoMatchRegexpCompiled)
}
