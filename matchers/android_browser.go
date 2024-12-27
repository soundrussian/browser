package matchers

import "strings"

type AndroidBrowser struct {
	p Parser
}

var (
	androidBrowserName          = "Android Browser"
	androidBrowserVersionRegexp = []string{`Version/([\d.]+)`}
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
	return b.p.Version(androidBrowserVersionRegexp, 1)
}

func (b *AndroidBrowser) Match() bool {
	// Without lookbehinds in Go regexp library, it is much easier to use strings.Contains
	ua := strings.ToLower(b.p.String())

	return strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android")
}
