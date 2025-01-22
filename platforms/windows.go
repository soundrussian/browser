package platforms

import "github.com/soundrussian/browser/v2/utils"

type Windows struct {
	p Parser
}

var (
	windowsName                  = "Windows"
	windowsVersionRegexp         = []string{`Windows NT\s*([0-9_.]+)?`}
	windowsMatchRegexp           = []string{`Windows`}
	windowsVersionRegexpCompiled = utils.CompileRegexps(windowsVersionRegexp)
	windowsMatchRegexpCompiled   = utils.CompileRegexps(windowsMatchRegexp)
)

func NewWindows(p Parser) *Windows {
	return &Windows{
		p: p,
	}
}

func (w *Windows) Name() string {
	return windowsName
}

func (w *Windows) Version() string {
	return w.p.Version(windowsVersionRegexpCompiled, 1, "0")
}

func (w *Windows) Match() bool {
	return w.p.Match(windowsMatchRegexpCompiled)
}
