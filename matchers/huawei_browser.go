package matchers

import "github.com/soundrussian/browser/v2/utils"

type HuaweiBrowser struct {
	p Parser
}

var (
	huaweiBrowserName                  = "Huawei Browser"
	huaweiBrowserVersionRegexp         = []string{`(?i)(?:HuaweiBrowser|HBPC)/([\d.]+)`}
	huaweiBrowserMatchRegexp           = []string{`(?i)(HuaweiBrowser|HBPC)`}
	huaweiBrowserVersionRegexpCompiled = utils.CompileRegexps(huaweiBrowserVersionRegexp)
	huaweiBrowserMatchRegexpCompiled   = utils.CompileRegexps(huaweiBrowserMatchRegexp)
)

func NewHuaweiBrowser(p Parser) *HuaweiBrowser {
	return &HuaweiBrowser{
		p: p,
	}
}

func (h *HuaweiBrowser) Name() string {
	return huaweiBrowserName
}

func (h *HuaweiBrowser) Version() string {
	return h.p.Version(huaweiBrowserVersionRegexpCompiled, 1)
}

func (h *HuaweiBrowser) Match() bool {
	return h.p.Match(huaweiBrowserMatchRegexpCompiled)
}
