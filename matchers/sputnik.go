package matchers

import "github.com/soundrussian/browser/v2/utils"

type Sputnik struct {
	p Parser
}

var (
	sputnikName                  = "Sputnik"
	sputnikVersionRegexp         = []string{`SputnikBrowser/([\d.]+)`}
	sputnikMatchRegexp           = []string{`SputnikBrowser`}
	sputnikVersionRegexpCompiled = utils.CompileRegexps(sputnikVersionRegexp)
	sputnikMatchRegexpCompiled   = utils.CompileRegexps(sputnikMatchRegexp)
)

func NewSputnik(p Parser) *Sputnik {
	return &Sputnik{
		p: p,
	}
}

func (s *Sputnik) Name() string {
	return sputnikName
}

func (s *Sputnik) Version() string {
	return s.p.Version(sputnikVersionRegexpCompiled, 1)
}

func (s *Sputnik) Match() bool {
	return s.p.Match(sputnikMatchRegexpCompiled)
}
