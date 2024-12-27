package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCocCoc(t *testing.T) {
	Convey("Subject: #NewCocCoc", t, func() {
		Convey("It should return a new CocCoc instance", func() {
			So(NewCocCoc(NewUAParser("")), ShouldHaveSameTypeAs, &CocCoc{})
		})
	})
}

func TestCocCocName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return Coc Coc", func() {
			sb := NewCocCoc(NewUAParser(""))
			So(sb.Name(), ShouldEqual, "Coc Coc")
		})
	})
}

func TestCocCocVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is matched", func() {
			Convey("It should return the version", func() {
				sb := testUserAgents["coc-coc"]

				So(NewCocCoc(NewUAParser(sb.Android)).Version(), ShouldEqual, "80.0.232")
				So(NewCocCoc(NewUAParser(sb.Mac)).Version(), ShouldEqual, "94.0.210")
				So(NewCocCoc(NewUAParser(sb.Windows)).Version(), ShouldEqual, "56.3.162")
			})
		})

		Convey("When the version is not matched", func() {
			Convey("It should return default version", func() {
				sb := NewCocCoc(NewUAParser("Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H)"))
				So(sb.Version(), ShouldEqual, "0.0")
			})
		})
	})
}

func TestCocCocMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Coc Coc", func() {
			Convey("It should return true", func() {
				sb := testUserAgents["coc-coc"]

				So(NewCocCoc(NewUAParser(sb.Android)).Match(), ShouldBeTrue)
				So(NewCocCoc(NewUAParser(sb.Mac)).Match(), ShouldBeTrue)
				So(NewCocCoc(NewUAParser(sb.Windows)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Coc Coc", func() {
			Convey("It should return false", func() {
				sb := NewCocCoc(NewUAParser(testUserAgents["chrome"].Android))
				So(sb.Match(), ShouldBeFalse)
			})
		})
	})
}
