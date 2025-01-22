package matchers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBaidu(t *testing.T) {
	Convey("Subject: #NewBaidu", t, func() {
		Convey("It should return a new Baidu instance", func() {
			baidu := NewBaidu(NewUAParser(""))
			So(baidu, ShouldHaveSameTypeAs, &Baidu{})
		})
	})
}

func TestBaiduName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("It should return the correct name", func() {
			baidu := NewBaidu(NewUAParser(""))
			So(baidu.Name(), ShouldEqual, "Baidu")
		})
	})
}

func TestBaiduVersion(t *testing.T) {
	Convey("Subject: #Version", t, func() {
		Convey("When the version is captured", func() {
			Convey("It should return the version", func() {
				ua := testUserAgents["baidu"]

				So(NewBaidu(NewUAParser(ua.IOS)).Version(), ShouldEqual, "13.81.1.10")
				So(NewBaidu(NewUAParser(ua.Android)).Version(), ShouldEqual, "13.81.1.10")
			})
		})

	})
}

func TestBaiduMatch(t *testing.T) {
	Convey("Subject: #Match", t, func() {
		Convey("When user agent matches Baidu", func() {
			Convey("It should return true", func() {
				ua := testUserAgents["baidu"]
				So(NewBaidu(NewUAParser(ua.IOS)).Match(), ShouldBeTrue)
				So(NewBaidu(NewUAParser(ua.Android)).Match(), ShouldBeTrue)
			})
		})

		Convey("When user agent does not match Baidu", func() {
			Convey("It should return false", func() {
				ua := "Mozilla/5.0 (Linux; Android 10; Redmi K30 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Mobile Safari/537.36"
				baidu := NewBaidu(NewUAParser(ua))
				So(baidu.Match(), ShouldBeFalse)
			})
		})
	})
}
