package matchers

type Baidu struct {
	p Parser
}

var (
	baiduName          = "Baidu"
	baiduVersionRegexp = []string{`baiduboxapp/([\d.]+)`}
	baiduMatchRegex    = []string{`(?i)baiduboxapp`}
)

func NewBaidu(p Parser) *Baidu {
	return &Baidu{
		p: p,
	}
}

func (a *Baidu) Name() string {
	return baiduName
}

func (a *Baidu) Version() string {
	return a.p.Version(baiduVersionRegexp, 1)
}

func (a *Baidu) Match() bool {
	return a.p.Match(baiduMatchRegex)
}
