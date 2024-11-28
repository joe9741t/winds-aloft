package winds

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWindsAloft(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WindsAloft Suite")
}

var _ = Describe("WindsAloft", func() {
	Describe("AddWind", func() {
		It("should add wind data correctly", func() {
			wa := &WindsAloft{}

			// api response for wind data looks like this
			// AXN 3411 3314-10 3315-13 3116-19 3221-31 3236-39 294251 294155 283555
			// MSP 3413 3314-09 3115-13 3016-19 3118-31 3035-39 304849 284554 274454

			wind := Wind{
				Airport: "AXN",
				Ft0:     "3411",
				Ft3000:  "3314-10",
				Ft6000:  "3315-13",
				Ft9000:  "3116-19",
				Ft12000: "3221-31",
			}
			wa.AddWind(wind)
			windMsp := Wind{
				Airport: "MSP",
				Ft0:     "3413",
				Ft3000:  "3314-09",
				Ft6000:  "3315-13",
				Ft9000:  "3116-19",
				Ft12000: "3221-31",
			}
			wa.AddWind(windMsp)

			Expect(len(wa.Winds)).To(Equal(2))
			Expect(wa.Winds["AXN"]).To(Equal(wind))
		})
	})
})
