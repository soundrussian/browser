package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAndroidBrowser(t *testing.T) {
	Convey("Subject: #NewAndroidBrowser", t, func() {
		Convey("It should return a new Android Browser instance", func() {
			So(NewAndroidBrowser(NewUAParser("")), ShouldHaveSameTypeAs, &AndroidBrowser{})
		})
	})
}

func TestAndroidBrowserName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return AndroidBrowser", func() {
			sb := NewAndroidBrowser(NewUAParser(""))
			So(sb.Name(), ShouldEqual, "Android Browser")
		})
	})
}

func TestAndroidBrowserVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				sb := testUserAgents["android-browser"]

				So(NewAndroidBrowser(NewUAParser(sb.Android)).Version(), ShouldEqual, "4.0")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				sb := NewAndroidBrowser(NewUAParser("Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H)"))
				So(sb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestAndroidBrowserMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Android Browser", func() {
			Convey("It should return true", func() {
				sb := testUserAgents["android-browser"]

				So(NewAndroidBrowser(NewUAParser(sb.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Android Browser", func() {
			Convey("It should return false", func() {
				sb := NewAndroidBrowser(NewUAParser(testUserAgents["chrome"].Android))
				So(sb.Match(), ShouldBeFalse)
			})
		})
	})
}
