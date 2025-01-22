package matchers

import (
	"strings"

	"github.com/soundrussian/browser/v2/utils"
)

type AndroidBrowser struct {
	p Parser
}

var (
	androidBrowserName                  = "Android Browser"
	androidBrowserVersionRegexp         = []string{`Version/([\d.]+)`}
	androidBrowserVersionRegexpCompiled = utils.CompileRegexps(androidBrowserVersionRegexp)
)

func NewAndroidBrowser(p Parser) *AndroidBrowser {
	return &AndroidBrowser{
		p: p,
	}
}

func (b *AndroidBrowser) Name() string {
	return androidBrowserName
}

func (b *AndroidBrowser) Version() string {
	return b.p.Version(androidBrowserVersionRegexpCompiled, 1)
}

func (b *AndroidBrowser) Match() bool {
	// Without lookbehinds in Go regexp library, it is much easier to use strings.Contains
	ua := strings.ToLower(b.p.String())

	return strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android")
}
