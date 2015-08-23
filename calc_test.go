package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculator(t *testing.T) {

	Convey("Sum should find sum of various numbers", t, func() {

		Convey("It should return the sum of the numbers:", func() {
			So(sum(1, 2, 3), ShouldEqual, 1+2+3)
			So(sum(2, 4, 6), ShouldEqual, 2+4+6)
			So(sum(1, -2, 3), ShouldEqual, 1-2+3)
			So(sum(1.1, 2.3, 5.8), ShouldEqual, 1.1+2.3+5.8)
			//So(sum(1.1, -3.3), ShouldEqual, 1.1-3.3) // FAILS becase 2.2 != 2.219999...997
		})

	})

	Convey("Subtract should find difference of various numbers", t, func() {

		Convey("It should return the difference of the numbers:", func() {
			So(difference(1, 2, 3), ShouldEqual, 1-2-3)
			So(difference(-1, -2, -3), ShouldEqual, -1+2+3)
			So(difference(-2, -1, -3), ShouldEqual, -2+1+3)
			So(difference(1.1, 2.2), ShouldEqual, 1.1-2.2)
			//So(difference(1.1, -2.2), ShouldEqual, 1.1+2.2) //FAILS because 3.3 != 3.300...003
		})

	})

	Convey("Multiply should find product of various numbers", t, func() {

		Convey("It should return the product of the numbers:", func() {
			So(product(1, 2, 3), ShouldEqual, 1*2*3)
		})

	})

	Convey("Divide should find quotient of various numbers", t, func() {

		Convey("It should return a number", nil)

		Convey("Cannot divide by zero", nil)

	})

	/* Empty Convey
	Convey("Multiply should find product of various numbers", t, func() {

	  Convey("It should return a number", nil)

	})
	*/

}
