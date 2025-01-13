package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewXbox(t *testing.T) {
	Convey("Subject: #NewXbox", t, func() {
		Convey("It should return a new Xbox instance", func() {
			So(NewXbox(NewUAParser("")), ShouldHaveSameTypeAs, &Xbox{})
		})
	})
}

func TestXboxName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Xbox", func() {
			So(NewXbox(NewUAParser("")).Name(), ShouldEqual, "Xbox")
		})
	})
}

func TestXboxVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It does not detect version", func() {
				So(NewXbox(NewUAParser(testPlatforms["xbox"])).Version(), ShouldEqual, "")
			})
		})

		Convey("When is not Xbox", func() {
			Convey("It should return default version", func() {
				So(NewXbox(NewUAParser(testPlatforms["firefox"])).Version(), ShouldEqual, "")
			})
		})
	})
}

func TestXboxMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Xbox", func() {
			Convey("It should return true", func() {
				So(NewXbox(NewUAParser(testPlatforms["xbox"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Xbox", func() {
			Convey("It should return false", func() {
				So(NewXbox(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
