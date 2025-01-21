package platforms

import "github.com/soundrussian/browser/v2/utils"

type Playstation struct {
	p Parser
}

var (
	playstationOSName                  = "Playstation"
	playstationOSVersionRegexp         = []string{`(?i)playstation(\s*(\d+|vita))?[ \/]?([\d.]+)`}
	playstationOSMatchRegexp           = []string{`(?i)playstation`}
	playstationOSVersionRegexpCompiled = utils.CompileRegexps(playstationOSVersionRegexp)
	playstationOSMatchRegexpCompiled   = utils.CompileRegexps(playstationOSMatchRegexp)
)

func NewPlaystation(p Parser) *Playstation {
	return &Playstation{
		p: p,
	}
}

func (w *Playstation) Name() string {
	return playstationOSName
}

func (w *Playstation) Version() string {
	globalVersion := w.p.Version(playstationOSVersionRegexpCompiled, 1, "0.0")
	if globalVersion == "0.0" || globalVersion == "" {
		return "0.0"
	}
	return w.p.Version(playstationOSVersionRegexpCompiled, 3, "0.0")
}

func (w *Playstation) Match() bool {
	return w.p.Match(playstationOSMatchRegexpCompiled)
}
