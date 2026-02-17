package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAlohaBrowser(t *testing.T) {
	Convey("Subject: #NewAlohaBrowser", t, func() {
		Convey("It should return a new AlohaBrowser instance", func() {
			So(NewAlohaBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &AlohaBrowser{})
		})
	})
}

func TestAlohaBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Aloha Browser", func() {
			sb := NewAlohaBrowser(NewUAParser(""))
			So(sb.Name(), ShouldEqual, "Aloha Browser")
		})
	})
}

func TestAlohaBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				sb := testUserAgents["aloha-browser"]

				So(NewAlohaBrowser(NewUAParser(sb.Android)).Version(), ShouldEqual, "5.10.0")
				So(NewAlohaBrowser(NewUAParser(sb.IOS)).Version(), ShouldEqual, "7.11.2")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				sb := NewAlohaBrowser(NewUAParser("Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H)"))
				So(sb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestAlohaBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Aloha Browser", func() {
			Convey("It should return true", func() {
				sb := testUserAgents["aloha-browser"]

				So(NewAlohaBrowser(NewUAParser(sb.Android)).Match(), ShouldBeTrue)
				So(NewAlohaBrowser(NewUAParser(sb.IOS)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Aloha Browser", func() {
			Convey("It should return false", func() {
				sb := NewAlohaBrowser(NewUAParser(testUserAgents["chrome"].Android))
				So(sb.Match(), ShouldBeFalse)
			})
		})
	})
}
