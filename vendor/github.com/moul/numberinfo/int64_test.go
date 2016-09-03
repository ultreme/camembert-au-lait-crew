package numberinfo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt64Number(t *testing.T) {
	Convey("Testing Int64Number", t, func() {
		_, ok := interface{}(Int64(4242)).(Number)
		So(ok, ShouldBeTrue)
	})
}

func TestInt64(t *testing.T) {
	Convey("Testing Int64()", t, func() {
		So(Int64(4242).value, ShouldEqual, 4242)
		So(Int64(-4242).value, ShouldEqual, -4242)
	})
}

func TestInt64Number_Float64(t *testing.T) {
	Convey("Testing Int64Number.Int64()", t, func() {
		val, err := Int64(4242).Float64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, 4242.0)

		val, err = Int64(-4242).Float64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, -4242.0)
	})
}

func TestInt64Number_Int64(t *testing.T) {
	Convey("Testing Int64Number.Int64()", t, func() {
		val, err := Int64(4242).Int64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, 4242)

		val, err = Int64(-4242).Int64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, -4242)
	})
}

func TestInt64Number_BigFactorial(t *testing.T) {
	Convey("Testing Int64Number.BigFactorial()", t, func() {
		val, err := Int64(5).BigFactorial()
		So(err, ShouldBeNil)
		So(val.Uint64(), ShouldEqual, 120)

		val, err = Int64(-5).BigFactorial()
		So(err, ShouldBeNil)
		So(val.Uint64(), ShouldEqual, 1)
	})
}
