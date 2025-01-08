package platforms

type Playstation struct {
	p Parser
}

var (
	playstationOSName          = "Playstation"
	playstationOSVersionRegexp = []string{`(?i)playstation(\s*(\d+|vita))?[ \/]?([\d.]+)`}
	playstationOSMatchRegexp   = []string{`(?i)playstation`}
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
	globalVersion := w.p.Version(playstationOSVersionRegexp, 1, "0.0")
	if globalVersion == "0.0" || globalVersion == "" {
		return "0.0"
	}
	return w.p.Version(playstationOSVersionRegexp, 3, "0.0")
}

func (w *Playstation) Match() bool {
	return w.p.Match(playstationOSMatchRegexp)
}
