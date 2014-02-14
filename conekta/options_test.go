package conekta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FilterOptions", func() {

	Context("Operators", func() {
		var f *FilterOptions

		BeforeEach(func() {
			f = NewFilterOptions()
		})

		It("Eq", func() {
			f.Eq("amount", 20000)
			Expect(f.Q()).To(Equal("?amount=20000"))
		})

		It("Gt", func() {
			f.Gt("amount", 20000)
			Expect(f.Q()).To(Equal("?amount.gt=20000"))
		})

		It("Gte", func() {
			f.Gte("amount", 20000)
			Expect(f.Q()).To(Equal("?amount.gte=20000"))
		})

		It("Lt", func() {
			f.Lt("amount", 20000)
			Expect(f.Q()).To(Equal("?amount.lt=20000"))

		})

		It("Lte", func() {
			f.Lte("amount", 20000)
			Expect(f.Q()).To(Equal("?amount.lte=20000"))

		})

		It("Ne", func() {
			f.Ne("amount", "paid")
			Expect(f.Q()).To(Equal("?amount.ne=paid"))
		})

		PIt("In", func() {})

		PIt("Nin", func() {})

		It("Regex", func() {
			f.Regex("description", "pancakes")
			Expect(f.Q()).To(Equal("?description.regex=pancakes"))
		})

		It("Limit", func() {
			f.Limit(1)
			Expect(f.Q()).To(Equal("?limit=1"))
		})

		It("Offset", func() {
			f.Offset(10)
			Expect(f.Q()).To(Equal("?offset=10"))
		})

		It("Sort", func() {
			f.Sort("amount", "desc")
			Expect(f.Q()).To(Equal("?sort=amount.desc"))
		})

	})

	Context("Append conditions", func() {
		f := NewFilterOptions()

		It("Combines operators", func() {
			f.Gt("amount", 20000).Regex("description", "pancakes").Limit(2).Offset(10).Sort("amount", "desc")
			Expect(f.Q()).To(Equal("?amount.gt=20000&description.regex=pancakes&limit=2&offset=10&sort=amount.desc"))
		})

	})
})
