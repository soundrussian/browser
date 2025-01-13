package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewNintendo(t *testing.T) {
	Convey("Subject: #NewNintendo", t, func() {
		Convey("It should return a new Nintendo instance", func() {
			So(NewNintendo(NewUAParser("")), ShouldHaveSameTypeAs, &Nintendo{})
		})
	})
}

func TestNintendoName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Nintendo", func() {
			So(NewNintendo(NewUAParser("")).Name(), ShouldEqual, "Nintendo")
		})
	})
}

func TestNintendoVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It does not detect version", func() {
				So(NewNintendo(NewUAParser(testPlatforms["nintendo"])).Version(), ShouldEqual, "")
			})
		})

		Convey("When is not Nintendo", func() {
			Convey("It should return default version", func() {
				So(NewNintendo(NewUAParser(testPlatforms["firefox"])).Version(), ShouldEqual, "")
			})
		})
	})
}

func TestNintendoMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Nintendo", func() {
			Convey("It should return true", func() {
				So(NewNintendo(NewUAParser(testPlatforms["nintendo"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Nintendo", func() {
			Convey("It should return false", func() {
				So(NewNintendo(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
