package matchers

import "github.com/soundrussian/browser/v2/utils"

type Electron struct {
	p Parser
}

var (
	electronName                  = "Electron"
	electronVersionRegexp         = []string{`Electron/([\d.]+)`}
	electronMatchRegex            = []string{`Electron`}
	electronVersionRegexpCompiled = utils.CompileRegexps(electronVersionRegexp)
	electronMatchRegexCompiled    = utils.CompileRegexps(electronMatchRegex)
)

func NewElectron(p Parser) *Electron {
	return &Electron{
		p: p,
	}
}

func (e *Electron) Name() string {
	return electronName
}

func (e *Electron) Version() string {
	return e.p.Version(electronVersionRegexpCompiled, 1)
}

func (e *Electron) Match() bool {
	return e.p.Match(electronMatchRegexCompiled)
}
