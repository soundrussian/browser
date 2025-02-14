package matchers

import "github.com/soundrussian/browser/v2/utils"

type CocCoc struct {
	p Parser
}

var (
	cocCocName         = "Coc Coc"
	cocCocVersionRegex = []string{
		`(?i)coc_coc_browser/([\d.]+)`,
		`Chrome/([\d.]+)`,
		`Safari/([\d.]+)`,
	}
	cocCocMatchRegex           = []string{`coc_coc_browser/`}
	cocCocVersionRegexCompiled = utils.CompileRegexps(cocCocVersionRegex)
	cocCocMatchRegexCompiled   = utils.CompileRegexps(cocCocMatchRegex)
)

func NewCocCoc(p Parser) *CocCoc {
	return &CocCoc{
		p: p,
	}
}

func (s *CocCoc) Name() string {
	return cocCocName
}

func (s *CocCoc) Version() string {
	return s.p.Version(cocCocVersionRegexCompiled, 1)
}

func (s *CocCoc) Match() bool {
	return s.p.Match(cocCocMatchRegexCompiled)
}
