package matchers

import "github.com/soundrussian/browser/v2/utils"

type Konqueror struct {
	p Parser
}

var (
	konquerorName                  = "Konqueror"
	konquerorVersionRegexp         = []string{`(?i)Konqueror/([\d.]+)`}
	konquerorMatchRegex            = []string{`(?i)Konqueror`}
	konquerorVersionRegexpCompiled = utils.CompileRegexps(konquerorVersionRegexp)
	konquerorMatchRegexCompiled    = utils.CompileRegexps(konquerorMatchRegex)
)

func NewKonqueror(p Parser) *Konqueror {
	return &Konqueror{
		p: p,
	}
}

func (k *Konqueror) Name() string {
	return konquerorName
}

func (k *Konqueror) Version() string {
	return k.p.Version(konquerorVersionRegexpCompiled, 1)
}

func (k *Konqueror) Match() bool {
	return k.p.Match(konquerorMatchRegexCompiled)
}
