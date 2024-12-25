package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSilkBrowser(t *testing.T) {
	Convey("Subject: #NewSilkBrowser", t, func() {
		Convey("It should return a new SilkBrowser instance", func() {
			So(NewSilkBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &SilkBrowser{})
		})
	})
}

func TestSilkBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Silk", func() {
			sb := NewSilkBrowser(NewUAParser(""))
			So(sb.Name(), ShouldEqual, "Silk")
		})
	})
}

func TestSilkBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				sb := testUserAgents["silk-browser"]

				So(NewSilkBrowser(NewUAParser(sb.Android)).Version(), ShouldEqual, "130.5.9")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				sb := NewSilkBrowser(NewUAParser("Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H)"))
				So(sb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestSilkBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Silk Browser", func() {
			Convey("It should return true", func() {
				sb := testUserAgents["silk-browser"]

				So(NewSilkBrowser(NewUAParser(sb.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Silk Browser", func() {
			Convey("It should return false", func() {
				sb := NewSilkBrowser(NewUAParser(testUserAgents["chrome"].Linux))
				So(sb.Match(), ShouldBeFalse)
			})
		})
	})
}
