package matchers

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
			n := NewNintendo(NewUAParser(""))
			So(n.Name(), ShouldEqual, "Nintendo")
		})
	})
}

func TestNintendoVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		ua := testUserAgents["nintendo"]
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				So(NewNintendo(NewUAParser(ua.Linux)).Version(), ShouldEqual, "5.1.0.23653")
			})
		})
	})

	Convey("When the version is not matched", t, func() {
		Convey("It should return default version", func() {
			n := NewNintendo(NewUAParser(testUserAgents["chrome"].Linux))
			So(n.Version(), ShouldEqual, "0.0")
		})
	})
}

func TestNintendoMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Nintendo", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["nintendo"]

				So(NewNintendo(NewUAParser(ua.Linux)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Nintendo", func() {
			Convey("It should return false", func() {
				ua := testUserAgents["chrome"]

				So(NewNintendo(NewUAParser(ua.Linux)).Match(), ShouldBeFalse)
			})
		})
	})
}
