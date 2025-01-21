package matchers

import "github.com/soundrussian/browser/v2/utils"

type Instagram struct {
	p Parser
}

var (
	instagramName                  = "Instagram"
	instagramVersionRegexp         = []string{`Instagram[ /]([\d.]+)`}
	instagramMatchRegex            = []string{`(?i)Instagram`}
	instagramVersionRegexpCompiled = utils.CompileRegexps(instagramVersionRegexp)
	instagramMatchRegexCompiled    = utils.CompileRegexps(instagramMatchRegex)
)

func NewInstagram(p Parser) *Instagram {
	return &Instagram{
		p: p,
	}
}

func (i *Instagram) Name() string {
	return instagramName
}

func (i *Instagram) Version() string {
	return i.p.Version(instagramVersionRegexpCompiled, 1)
}

func (i *Instagram) Match() bool {
	return i.p.Match(instagramMatchRegexCompiled)
}
