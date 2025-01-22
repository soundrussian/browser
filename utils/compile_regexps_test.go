package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCompileRegexps(t *testing.T) {
	Convey("Subject: #CompileRegexps", t, func() {
		Convey("It compiles a given list of strings as regexps", func() {
			list := []string{`(?:Mac|MAC) OS X\s*([0-9_.]+)?`, `iPad`}
			actual := CompileRegexps(list)
			So(actual, ShouldNotBeNil)
			So(actual, ShouldHaveLength, 2)
			So(actual[0].String(), ShouldEqual, `(?:Mac|MAC) OS X\s*([0-9_.]+)?`)
			So(actual[1].String(), ShouldEqual, `iPad`)
		})
	})
}
