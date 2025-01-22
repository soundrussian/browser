package devices

import "github.com/soundrussian/browser/v2/utils"

type WiiU struct {
	p Parser
}

var (
	wiiuName               = "Nintendo WiiU"
	wiiUMatchRegex         = []string{`(?i)Nintendo WiiU`}
	wiiUMatchRegexCompiled = utils.CompileRegexps(wiiUMatchRegex)
)

func NewWiiU(p Parser) *WiiU {
	return &WiiU{
		p: p,
	}
}

func (w *WiiU) Name() string {
	return wiiuName
}

func (w *WiiU) Match() bool {
	return w.p.Match(wiiUMatchRegexCompiled)
}
