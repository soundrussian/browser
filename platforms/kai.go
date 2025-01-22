package platforms

import "github.com/soundrussian/browser/v2/utils"

type KaiOS struct {
	p Parser
}

var (
	kaiOSName                  = "Kai OS"
	kaiOSVersionRegexp         = []string{`KaiOS/([\d.]+)`}
	kaiOSMatchRegexp           = []string{`KaiOS`}
	kaiOSVersionRegexpCompiled = utils.CompileRegexps(kaiOSVersionRegexp)
	kaiOSMatchRegexpCompiled   = utils.CompileRegexps(kaiOSMatchRegexp)
)

func NewKaiOS(p Parser) *KaiOS {
	return &KaiOS{
		p: p,
	}
}

func (k *KaiOS) Name() string {
	return kaiOSName
}

func (k *KaiOS) Version() string {
	return k.p.Version(kaiOSVersionRegexpCompiled, 1, "")
}

func (k *KaiOS) Match() bool {
	return k.p.Match(kaiOSMatchRegexpCompiled)
}
