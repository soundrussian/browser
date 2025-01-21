package utils

import "regexp"

// CompileRegexps compiles a slice of string patterns into a slice of regular expressions using regexp.MustCompile.
func CompileRegexps(patterns []string) []*regexp.Regexp {
	result := make([]*regexp.Regexp, len(patterns))

	for i, pattern := range patterns {
		result[i] = regexp.MustCompile(pattern)
	}

	return result
}
