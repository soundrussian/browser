package devices

import (
	"regexp"

	"github.com/soundrussian/browser/v2/utils"
)

type Android struct {
	p Parser
}

var (
	androidNameRegex          = `(?i)\(Linux.*?; Android.*?; ([-_a-z0-9 ]+)(?:;)? Build[^)]+\)`
	androidMatchRegex         = []string{`(?i)Android`}
	androidNameRegexCompiled  = regexp.MustCompile(androidNameRegex)
	androidMatchRegexCompiled = utils.CompileRegexps(androidMatchRegex)
)

func NewAndroid(p Parser) *Android {
	return &Android{
		p: p,
	}
}

func (a Android) Name() string {
	matches := androidNameRegexCompiled.FindStringSubmatch(a.p.String())
	if len(matches) > 1 {
		return matches[1]
	}

	return "Unknown"
}

func (a Android) Match() bool {
	return a.p.Match(androidMatchRegexCompiled)
}
