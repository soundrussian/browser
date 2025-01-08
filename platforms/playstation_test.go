package platforms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPlaystation(t *testing.T) {
	Convey("Subject: #NewPlaystation", t, func() {
		Convey("It should return a new Playstation instance", func() {
			So(NewPlaystation(NewUAParser("")), ShouldHaveSameTypeAs, &Playstation{})
		})
	})
}

func TestPlaystationName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Playstation", func() {
			So(NewPlaystation(NewUAParser(testPlatforms["playstation-3"])).Name(), ShouldEqual, "Playstation")
			So(NewPlaystation(NewUAParser(testPlatforms["playstation-4"])).Name(), ShouldEqual, "Playstation")
			So(NewPlaystation(NewUAParser(testPlatforms["playstation-vita"])).Name(), ShouldEqual, "Playstation")
		})
	})
}

func TestPlaystationVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-3"])).Version(), ShouldEqual, "4.91")
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-4"])).Version(), ShouldEqual, "12.00")
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-5"])).Version(), ShouldEqual, "10.40")
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-vita"])).Version(), ShouldEqual, "3.74")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-smarttv"])).Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestPlaystationMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Playstation", func() {
			Convey("It should return true", func() {
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-3"])).Match(), ShouldBeTrue)
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-4"])).Match(), ShouldBeTrue)
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-5"])).Match(), ShouldBeTrue)
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-smarttv"])).Match(), ShouldBeTrue)
				So(NewPlaystation(NewUAParser(testPlatforms["playstation-vita"])).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Playstation", func() {
			Convey("It should return false", func() {
				So(NewPlaystation(NewUAParser(testPlatforms["firefox"])).Match(), ShouldBeFalse)
			})
		})
	})
}
