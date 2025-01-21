package matchers

import "github.com/soundrussian/browser/v2/utils"

type Opera struct {
	p Parser
}

var (
	operaName = "Opera"
	// order matters for version detection
	operaVersionRegexp         = []string{`Opera Mini/([\d.]+)`, `OP(?:R|iOS|T)/([\d.]+)`, `Opera/([\d.]+)`, `Version/([\d.]+)`}
	operaMatchRegexp           = []string{`(Opera|OP(R|iOS|T)/)`}
	operaVersionRegexpCompiled = utils.CompileRegexps(operaVersionRegexp)
	operaMatchRegexpCompiled   = utils.CompileRegexps(operaMatchRegexp)
)

func NewOpera(p Parser) *Opera {
	return &Opera{
		p: p,
	}
}

func (o *Opera) Name() string {
	return operaName
}

func (o *Opera) Version() string {
	return o.p.Version(operaVersionRegexpCompiled, 1)
}

func (o *Opera) Match() bool {
	return o.p.Match(operaMatchRegexpCompiled)
}
