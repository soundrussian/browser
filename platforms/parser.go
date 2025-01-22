package platforms

import "regexp"

type Parser interface {
	Version([]*regexp.Regexp, int, string) string
	Match([]*regexp.Regexp) bool
	String() string
}

type UAParser struct {
	userAgent string
}

// NewUAParser returns a new UAParser.
func NewUAParser(userAgent string) *UAParser {
	return &UAParser{
		userAgent: userAgent,
	}
}

func (b UAParser) String() string {
	return b.userAgent
}

// Version returns the version of the platform.
// The pattern is a list of compiled regular expressions.
// The version is extracted from the user agent string using the given patterns.
// The order parameter specifies which match to return.
func (b UAParser) Version(patterns []*regexp.Regexp, order int, defaultVersion string) string {
	for _, pattern := range patterns {
		matches := pattern.FindStringSubmatch(b.userAgent)
		if len(matches) > order {
			return matches[order]
		}
	}
	return defaultVersion
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
