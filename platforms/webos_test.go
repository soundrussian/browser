package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWebOS(t *testing.T) {
	Convey("Subject: #NewWebOS", t, func() {
		Convey("It should return a new WebOS instance", func() {
			So(NewWebOS(NewUAParser("")), ShouldHaveSameTypeAs, &WebOS{})
		})
	})
}

func TestWebOSName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return WebOS", func() {
			So(NewWebOS(NewUAParser("")).Name(), ShouldEqual, "WebOS")
		})
	})
}

func TestWebOSVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewWebOS(NewUAParser(testPlatforms["webos"])).Version(), ShouldEqual, "4.1.0")
				So(NewWebOS(NewUAParser(testPlatforms["webos-hp"])).Version(), ShouldEqual, "3.0.5")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewWebOS(NewUAParser(testPlatforms["firefox"])).Version(), ShouldEqual, "")
			})
		})
	})
}

func TestWebOSMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches WebOS", func() {
			Convey("It should return true", func() {
				So(NewWebOS(NewUAParser(testPlatforms["webos"])).Match(), ShouldBeTrue)
				So(NewWebOS(NewUAParser(testPlatforms["webos-hp"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match WebOS", func() {
			Convey("It should return false", func() {
				So(NewWebOS(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
