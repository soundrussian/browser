package matchers

import "github.com/soundrussian/browser/v2/utils"

type Snapchat struct {
	p Parser
}

var (
	snapchatName                  = "Snapchat"
	snapchatVersionRegexp         = []string{`Snapchat(?: |/)?([\d.]+)`}
	snapchatMatchRegexp           = []string{`Snapchat`}
	snapchatVersionRegexpCompiled = utils.CompileRegexps(snapchatVersionRegexp)
	snapchatMatchRegexpCompiled   = utils.CompileRegexps(snapchatMatchRegexp)
)

func NewSnapchat(p Parser) *Snapchat {
	return &Snapchat{
		p: p,
	}
}

func (s *Snapchat) Name() string {
	return snapchatName
}

func (s *Snapchat) Version() string {
	return s.p.Version(snapchatVersionRegexpCompiled, 1)
}

func (s *Snapchat) Match() bool {
	return s.p.Match(snapchatMatchRegexpCompiled)
}
