package matchers

import "github.com/soundrussian/browser/v2/utils"

type MicroMessenger struct {
	p Parser
}

var (
	microMessengerName                  = "MicroMessenger"
	microMessengerVersionRegexp         = []string{`(?i)(?:MicroMessenger)/([\d.]+)`}
	microMessengerMatchRegexp           = []string{`MicroMessenger`}
	microMessengerVersionRegexpCompiled = utils.CompileRegexps(microMessengerVersionRegexp)
	microMessengerMatchRegexpCompiled   = utils.CompileRegexps(microMessengerMatchRegexp)
)

func NewMicroMessenger(p Parser) *MicroMessenger {
	return &MicroMessenger{
		p: p,
	}
}

func (m *MicroMessenger) Name() string {
	return microMessengerName
}

func (m *MicroMessenger) Version() string {
	return m.p.Version(microMessengerVersionRegexpCompiled, 1)
}

func (m *MicroMessenger) Match() bool {
	return m.p.Match(microMessengerMatchRegexpCompiled)
}
