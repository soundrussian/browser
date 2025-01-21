package platforms

import "github.com/soundrussian/browser/v2/utils"

type Linux struct {
	p Parser
}

var (
	linuxName                = "Generic Linux"
	linuxMatchRegexp         = []string{`Linux`}
	linuxMatchRegexpCompiled = utils.CompileRegexps(linuxMatchRegexp)
)

func NewLinux(p Parser) *Linux {
	return &Linux{
		p: p,
	}
}

func (l *Linux) Name() string {
	return linuxName
}

func (l *Linux) Version() string {
	return "0"
}

func (l *Linux) Match() bool {
	return l.p.Match(linuxMatchRegexpCompiled)
}
