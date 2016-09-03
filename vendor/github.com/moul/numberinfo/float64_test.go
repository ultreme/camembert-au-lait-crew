package numberinfo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFloat64Number(t *testing.T) {
	Convey("Testing Float64Number", t, func() {
		_, ok := interface{}(Float64(4242)).(Number)
		So(ok, ShouldBeTrue)
	})
}

func TestFloat64(t *testing.T) {
	Convey("Testing Float64()", t, func() {
		So(Float64(4242.0).value, ShouldEqual, 4242)
		So(Float64(-4242.0).value, ShouldEqual, -4242)
	})
}

func TestFloat64Number_Float64(t *testing.T) {
	Convey("Testing Float64Number.Float64Number()", t, func() {
		val, err := Float64(4242.0).Float64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, 4242.0)

		val, err = Float64(-4242.0).Float64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, -4242.0)
	})
}

func TestFloat64Number_Int64(t *testing.T) {
	Convey("Testing Float64Number.Int64()", t, func() {
		val, err := Float64(4242.0).Int64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, 4242)

		val, err = Float64(-4242.0).Int64Number()
		So(err, ShouldBeNil)
		So(val.value, ShouldEqual, -4242)
	})
}

func TestFloat64Number_BigFactorial(t *testing.T) {
	Convey("Testing Float64Number.BigFactorial()", t, func() {
		val, err := Float64(5).BigFactorial()
		So(err, ShouldBeNil)
		So(val.Uint64(), ShouldEqual, 120)
	})
}
