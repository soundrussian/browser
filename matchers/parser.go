package matchers

import (
	"regexp"
)

// Parser is an interface for user agent parsers.
type Parser interface {
	String() string
	Version([]*regexp.Regexp, int) string
	Match([]*regexp.Regexp) bool
}

// compile time check if UAParser implements Parser interface
var _ Parser = (*UAParser)(nil)

type UAParser struct {
	userAgent string
}

func NewUAParser(userAgent string) *UAParser {
	return &UAParser{
		userAgent: userAgent,
	}
}

func (b UAParser) String() string {
	return b.userAgent
}

// Version returns the first match of the given patterns.
// The pattern is a list of regular expressions.
// The order is the index of the match group in the regular expression.
// If the order is greater than the number of matches, it returns "0.0".
func (b UAParser) Version(patterns []*regexp.Regexp, order int) string {
	for _, pattern := range patterns {
		matches := pattern.FindStringSubmatch(b.userAgent)

		if len(matches) > order {
			return matches[order]
		}
	}
	return "0.0"
}

// Match returns true if the user agent matches the pattern.
// The pattern is a list of regular expressions.
func (b UAParser) Match(patterns []*regexp.Regexp) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(b.userAgent) {
			return true
		}
	}
	return false
}
