package platforms

import (
	"strings"

	"github.com/soundrussian/browser/v2/utils"
)

type Android struct {
	p Parser
}

var (
	androidName                  = "Android"
	androidVersionRegexp         = []string{`(?i)Android ([\d.]+)`}
	androidMatchRegexp           = []string{`(?i)Android`}
	androidVersionRegexpCompiled = utils.CompileRegexps(androidVersionRegexp)
	androidMatchRegexpCompiled   = utils.CompileRegexps(androidMatchRegexp)
)

func NewAndroid(p Parser) *Android {
	return &Android{
		p: p,
	}
}

func (a *Android) Name() string {
	return androidName
}

func (a *Android) Version() string {
	return a.p.Version(androidVersionRegexpCompiled, 1, "")
}

func (a *Android) Match() bool {
	return a.p.Match(androidMatchRegexpCompiled) && !strings.Contains(a.p.String(), "KAIOS")
}
