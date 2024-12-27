package matchers

type Nintendo struct {
	p Parser
}

var (
	nintendoName          = "Nintendo"
	nintendoVersionRegexp = []string{`NintendoBrowser/([\d.]+)`}
	nintendoMatchRegexp   = []string{`NintendoBrowser`}
)

func NewNintendo(p Parser) *Nintendo {
	return &Nintendo{
		p: p,
	}
}

func (n *Nintendo) Name() string {
	return nintendoName
}

func (n *Nintendo) Version() string {
	return n.p.Version(nintendoVersionRegexp, 1)
}

func (n *Nintendo) Match() bool {
	return n.p.Match(nintendoMatchRegexp)
}
