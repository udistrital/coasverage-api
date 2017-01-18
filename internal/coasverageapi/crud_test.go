package coasverageapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDeleteCoverage(t *testing.T) {

	Convey("Given any id", t, func() {
		id := "any"
		Convey("It should Err", func() {
			So(DeleteCoverage(id), ShouldNotBeNil)
		})

	})

}
