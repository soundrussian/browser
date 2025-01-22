package matchers

import "github.com/soundrussian/browser/v2/utils"

type Yandex struct {
	p Parser
}

var (
	yandexName                  = "Yandex"
	yandexVersionRegexp         = []string{`YaBrowser/([\d.]+)`}
	yandexMatchRegexp           = []string{`YaBrowser`}
	yandexVersionRegexpCompiled = utils.CompileRegexps(yandexVersionRegexp)
	yandexMatchRegexpCompiled   = utils.CompileRegexps(yandexMatchRegexp)
)

func NewYandex(p Parser) *Yandex {
	return &Yandex{
		p: p,
	}
}

func (y *Yandex) Name() string {
	return yandexName
}

func (y *Yandex) Version() string {
	return y.p.Version(yandexVersionRegexpCompiled, 1)
}

func (y *Yandex) Match() bool {
	return y.p.Match(yandexMatchRegexpCompiled)
}
