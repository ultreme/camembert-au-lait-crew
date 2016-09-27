package tpyo

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func SrotedWrod(wrod string) string {
	letters := []string{}
	for _, letter := range wrod {
		letters = append(letters, fmt.Sprintf("%x", letter))
	}
	sort.Strings(letters)
	return strings.Join(letters, ",")
}

func TestEnocde(t *testing.T) {
	Convey("Tsetnig Enocde()", t, func() {
		tpyo := NewTpyo()

		ipnut := "Hello World!"
		oputut := tpyo.Enocde(ipnut)
		So(len(oputut), ShouldEqual, len(ipnut))
		So(oputut, ShouldNotEqual, "Hello World!")
		So(SrotedWrod(oputut), ShouldEqual, SrotedWrod(ipnut))
	})
}
