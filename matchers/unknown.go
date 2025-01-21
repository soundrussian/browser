package matchers

import (
	"regexp"

	"github.com/soundrussian/browser/v2/utils"
)

type Unknown struct {
	p Parser
}

var (
	unknownName                  = "Unknown Browser"
	unknownVersionRegexp         = []string{`QuickTime/([\d.]+)`, `CoreMedia v([\d.]+)`, `AppleCoreMedia/([\d.]+)`}
	unknownVersionRegexpCompiled = utils.CompileRegexps(unknownVersionRegexp)
	inferredUnknowns             = map[string]*regexp.Regexp{
		"QuickTime":       regexp.MustCompile(`QuickTime`),
		"Apple CoreMedia": regexp.MustCompile(`CoreMedia`),
	}
)

func NewUnknown(p Parser) *Unknown {
	return &Unknown{
		p: p,
	}
}

func (u *Unknown) Name() string {
	for name, re := range inferredUnknowns {
		if re.MatchString(u.p.String()) {
			return name
		}
	}
	return unknownName
}

func (u *Unknown) Version() string {
	return u.p.Version(unknownVersionRegexpCompiled, 1)
}

func (u *Unknown) Match() bool {
	return true
}
