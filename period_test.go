package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Periods", func() {
	const (
		testPeriodCodeJan uint32 = 201401
		testPeriodCodeFeb uint32 = 201402
		testPeriodCode501 uint32 = 201501
		testPeriodCode402 uint32 = 201402
	)

	Describe("Period", func() {
		It("should Compare_Different_Periods_AsEqual_When_2014_01", func() {
			testPeriodOne := NewPeriod(testPeriodCodeJan)

			testPeriodTwo := NewPeriod(testPeriodCodeJan)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeTrue())
		})

		It("should Compare_Different_Periods_AsEqual_When_2014_02", func() {
			testPeriodOne := NewPeriod(testPeriodCodeFeb)

			testPeriodTwo := NewPeriod(testPeriodCodeFeb)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeTrue())
		})

		It("should Compare_Different_Periods_SameYear_AsGreater", func() {
			testPeriodOne := NewPeriod(testPeriodCodeJan)

			testPeriodTwo := NewPeriod(testPeriodCodeFeb)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeFalse())

			Expect(testPeriodOne.OpLt(*testPeriodTwo)).To(BeTrue())
		})
		It("should Compare_Different_Periods_SameYear_AsLess", func() {
			testPeriodOne := NewPeriod(testPeriodCodeJan)

			testPeriodTwo := NewPeriod(testPeriodCodeFeb)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeFalse())

			Expect(testPeriodTwo.OpGt(*testPeriodOne)).To(BeTrue())
		})
		It("should Compare_Different_Periods_SameMonth_AsGreater", func() {
			testPeriodOne := NewPeriod(testPeriodCodeJan)

			testPeriodTwo := NewPeriod(testPeriodCodeFeb)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeFalse())

			Expect(testPeriodOne.OpLt(*testPeriodTwo)).To(BeTrue())
		})
		It("should Compare_Different_Periods_SameMonth_AsLess", func() {
			testPeriodOne := NewPeriod(testPeriodCodeJan)

			testPeriodTwo := NewPeriod(testPeriodCodeFeb)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeFalse())

			Expect(testPeriodTwo.OpGt(*testPeriodOne)).To(BeTrue())
		})
		It("should Compare_Different_Periods_DifferentYear_AsGreater", func() {
			testPeriodOne := NewPeriod(testPeriodCode402)

			testPeriodTwo := NewPeriod(testPeriodCode501)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeFalse())

			Expect(testPeriodOne.OpLt(*testPeriodTwo)).To(BeTrue())
		})
		It("should Compare_Different_Periods_DifferentYear_AsLess", func() {
			testPeriodOne := NewPeriod(testPeriodCode402)

			testPeriodTwo := NewPeriod(testPeriodCode501)

			Expect(testPeriodOne.Equals(*testPeriodTwo)).To(BeFalse())

			Expect(testPeriodTwo.OpGt(*testPeriodOne)).To(BeTrue())
		})
		It("should Return_Periods_Year_And_Month_2014_01", func() {
			testPeriodOne := NewPeriod(testPeriodCodeJan)

			Expect(testPeriodOne.Year()).To(Equal(2014))
			Expect(testPeriodOne.Month()).To(Equal(1))
		})
		It("should Return_PeriReturn_Periods_Year_And_Month_2014_02", func() {
			testPeriodOne := NewPeriod(testPeriodCodeFeb)

			Expect(testPeriodOne.Year()).To(Equal(2014))
			Expect(testPeriodOne.Month()).To(Equal(2))
		})
		It("should Return_Periods_Month_And_Year_Descriptions", func() {
			testPeriodJan := NewPeriod(testPeriodCodeJan)
			testPeriodFeb := NewPeriod(testPeriodCodeFeb)
			testPeriod501 := NewPeriod(testPeriodCode501)
			testPeriod402 := NewPeriod(testPeriodCode402)

			Expect(testPeriodJan.Description()).To(Equal("January 2014"))
			Expect(testPeriodFeb.Description()).To(Equal("February 2014"))
			Expect(testPeriod501.Description()).To(Equal("January 2015"))
			Expect(testPeriod402.Description()).To(Equal("February 2014"))
		})
	})
})
