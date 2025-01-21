package platforms

import "github.com/soundrussian/browser/v2/utils"

type WindowsMobile struct {
	p Parser
}

var (
	windowsMobileName                = "Windows Mobile"
	windowsMobileMatchRegexp         = []string{`Windows CE`}
	windowsMobileMatchRegexpCompiled = utils.CompileRegexps(windowsMobileMatchRegexp)
)

func NewWindowsMobile(p Parser) *WindowsMobile {
	return &WindowsMobile{
		p: p,
	}
}

func (w *WindowsMobile) Name() string {
	return windowsMobileName
}

func (w *WindowsMobile) Version() string {
	return "0"
}

func (w *WindowsMobile) Match() bool {
	return w.p.Match(windowsMobileMatchRegexpCompiled)
}
