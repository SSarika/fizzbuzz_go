package fb

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

/*func TestNumber(t *testing.T) {
	Convey("Given any number from 1 to 100, if we ask for a number, it returns that number", t, func() {
		So(fizzbuzz(11), ShouldEqual,11)
		} )
}*/

func TestNonMultipleNumber(t *testing.T) {
	Convey("Given any number from 1 to 100, if number is not a multiple of 3 and 5, it returns that number", t, func() {
	So(Fizzbuzz(7), ShouldEqual,"7")
} )
}

func Test3MultipleNumber(t *testing.T) {
	Convey("Given any number from 1 to 100, if number is a multiple of 3, it returns Fizz", t, func() {
		So(Fizzbuzz(3), ShouldEqual,"Fizz")
	} )
}

func Test5MultipleNumber(t *testing.T) {
	Convey("Given any number from 1 to 100, if number is a multiple of 5, it returns Buzz", t, func() {
		So(Fizzbuzz(10), ShouldEqual,"Buzz")
	} )
}

func Test3and5MultipleNumber(t *testing.T) {
	Convey("Given any number from 1 to 100, if number is a multiple of 3 and 5, it returns FizzBuzz", t, func() {
		So(Fizzbuzz(15), ShouldEqual,"FizzBuzz")
	} )
}
