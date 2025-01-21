package platforms

import "github.com/soundrussian/browser/v2/utils"

type Kindle struct {
	p Parser
}

var (
	kindleName               = "Kindle"
	kindleMatchRegex         = []string{`Kindle`}
	kindleMatchRegexCompiled = utils.CompileRegexps(kindleMatchRegex)
)

func NewKindle(p Parser) *Kindle {
	return &Kindle{
		p: p,
	}
}

func (k *Kindle) Name() string {
	return kindleName
}

// Version returns empty string for Kindle
func (k *Kindle) Version() string {
	return ""
}

func (k *Kindle) Match() bool {
	return k.p.Match(kindleMatchRegexCompiled)
}
