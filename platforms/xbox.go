package platforms

type Xbox struct {
	p Parser
}

var (
	xboxName       = "Xbox"
	xboxMatchRegex = []string{`Xbox`}
)

func NewXbox(p Parser) *Xbox {
	return &Xbox{
		p: p,
	}
}

func (k *Xbox) Name() string {
	return xboxName
}

// Version returns empty string for Xbox
func (k *Xbox) Version() string {
	return ""
}

func (k *Xbox) Match() bool {
	return k.p.Match(xboxMatchRegex)
}
