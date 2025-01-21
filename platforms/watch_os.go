package platforms

import "github.com/soundrussian/browser/v2/utils"

type WatchOS struct {
	p Parser
}

var (
	watchOSName                  = "Apple Watch OS"
	watchOSVersionRegexp         = []string{`(?i)Watch\s*OS[ ,/]([\d.]+)`, `Watch[^/]+/([\d.]+)`}
	watchOSMatchRegexp           = []string{`(?i)Watch\s*OS`, `Watch[\d+]`}
	watchOSVersionRegexpCompiled = utils.CompileRegexps(watchOSVersionRegexp)
	watchOSMatchRegexpCompiled   = utils.CompileRegexps(watchOSMatchRegexp)
)

func NewWatchOS(p Parser) *WatchOS {
	return &WatchOS{
		p: p,
	}
}

func (w *WatchOS) Name() string {
	return watchOSName
}

func (w *WatchOS) Version() string {
	return w.p.Version(watchOSVersionRegexpCompiled, 1, "0.0")
}

func (w *WatchOS) Match() bool {
	return w.p.Match(watchOSMatchRegexpCompiled)
}
