package devices

import "github.com/soundrussian/browser/v2/utils"

type KindleFire struct {
	p Parser
}

var (
	kindleFireName               = "Kindle Fire"
	kindleFireMatchRegex         = []string{`Kindle Fire|KFTT`}
	kindleFireMatchRegexCompiled = utils.CompileRegexps(kindleFireMatchRegex)
)

func NewKindleFire(p Parser) *KindleFire {
	return &KindleFire{
		p: p,
	}
}

func (k *KindleFire) Name() string {
	return kindleFireName
}

func (k *KindleFire) Match() bool {
	return k.p.Match(kindleFireMatchRegexCompiled)
}
