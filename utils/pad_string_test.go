package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPadStringWithSpaces(t *testing.T) {
	Convey("Subject: #PadStringWithSpaces", t, func() {
		Convey("It returns long enough string unchanged", func() {
			str := "test"
			So(PadStringWithSpaces(str, 4), ShouldEqual, "test")
		})

		Convey("It pads short string with spaces", func() {
			str := "t"
			So(PadStringWithSpaces(str, 4), ShouldEqual, "t   ")
		})
	})
}
