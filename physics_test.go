package physics_test

import (
	. "github.com/chelseatroy/physics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Physics", func() {
  Describe("Calculates Force", func() {
        Context("With mass and acceleration", func() {
            It("should calculate force", func() {
                Expect(Force(2.5, 4.0)).To(Equal(10.0))
            })
        })     
  })
})
