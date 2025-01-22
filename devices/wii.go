package devices

import "github.com/soundrussian/browser/v2/utils"

type Wii struct {
	p Parser
}

var (
	wiiName               = "Nintendo Wii"
	wiiMatchRegex         = []string{`(?i)Nintendo Wii`}
	wiiMatchRegexCompiled = utils.CompileRegexps(wiiMatchRegex)
)

func NewWii(p Parser) *Wii {
	return &Wii{
		p: p,
	}
}

func (w *Wii) Name() string {
	return wiiName
}

func (w *Wii) Match() bool {
	return w.p.Match(wiiMatchRegexCompiled)
}
