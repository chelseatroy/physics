package physics_test

import (
	. "github.com/chelseatroy/physics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Physics", func() {
  Describe("Force", func() {
        Context("With mass and acceleration", func() {
            It("should calculate force", func() {
                Expect(Force(2.5, 4.0)).To(Equal(10.0))
            })
        })     
  })

   Describe("Pendulums", func() {
        Context("With length of pendulum", func() {
            It("should calculate period", func() {
                Expect(PendulumPeriod(1.0)).To(Equal(2.007))
            })
        })     
  })
})
