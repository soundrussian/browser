package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewKindle(t *testing.T) {
	Convey("Subject: #NewKindle", t, func() {
		Convey("It should return a new Kindle instance", func() {
			So(NewKindle(NewUAParser("")), ShouldHaveSameTypeAs, &Kindle{})
		})
	})
}

func TestKindleName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Kindle", func() {
			So(NewKindle(NewUAParser("")).Name(), ShouldEqual, "Kindle")
		})
	})
}

func TestKindleVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It does not detect version", func() {
				So(NewKindle(NewUAParser(testPlatforms["kindle"])).Version(), ShouldEqual, "")
			})
		})

		Convey("When is not Kindle", func() {
			Convey("It should return default version", func() {
				So(NewKindle(NewUAParser(testPlatforms["firefox"])).Version(), ShouldEqual, "")
			})
		})
	})
}

func TestKindleMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Kindle", func() {
			Convey("It should return true", func() {
				So(NewKindle(NewUAParser(testPlatforms["kindle"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Kindle", func() {
			Convey("It should return false", func() {
				So(NewKindle(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
