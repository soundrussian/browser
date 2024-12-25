package matchers

type SilkBrowser struct {
	p Parser
}

var (
	silkBrowserName         = "Silk"
	silkBrowserVersionRegex = []string{
		`(?i)Silk/([\d.]+)`,
		`Chrome/([\d.]+)`,
		`Safari/([\d.]+)`,
	}
	silkBrowserMatchRegex = []string{`Silk/`}
)

func NewSilkBrowser(p Parser) *SilkBrowser {
	return &SilkBrowser{
		p: p,
	}
}

func (s *SilkBrowser) Name() string {
	return silkBrowserName
}

func (s *SilkBrowser) Version() string {
	return s.p.Version(silkBrowserVersionRegex, 1)
}

func (s *SilkBrowser) Match() bool {
	return s.p.Match(silkBrowserMatchRegex)
}
