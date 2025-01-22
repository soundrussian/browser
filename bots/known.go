package bots

import (
	"regexp"
)

// Known is a bot that matches when the user agent contains a keyword.
// It is used for bots that are known to be bots
type Known struct {
	userAgent  string
	bots       map[string]BotMatchInfo
	matchedBot string
}

type BotMatchInfo struct {
	Name   string
	Regexp *regexp.Regexp
}

// NewKnown returns a new Known bot.
func NewKnown(userAgent string, bots map[string]BotMatchInfo) *Known {
	return &Known{
		userAgent: userAgent,
		bots:      bots,
	}
}

// Name returns the keyword that matched.
func (k *Known) Name() string {
	if !k.matched() {
		k.Match()
	}

	return k.matchedBot
}

// Match returns true if the user agent contains a keyword.
func (k *Known) Match() bool {
	if k.matched() {
		return true
	}

	for _, n := range k.bots {
		if n.Regexp.MatchString(k.userAgent) {
			k.matchedBot = n.Name
			return true
		}
	}
	return false
}

func (k *Known) matched() bool {
	return k.matchedBot != ""
}

func CompileKnownBots(data map[string]string) map[string]BotMatchInfo {
	res := map[string]BotMatchInfo{}

	for bot, name := range data {
		res[bot] = BotMatchInfo{
			Name:   name,
			Regexp: regexp.MustCompile(`(?i)` + bot),
		}
	}

	return res
}
