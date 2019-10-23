package Importers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Augora/Augora-GraphQL/Importers"
	"github.com/Augora/Augora-GraphQL/Models"
)

var _ = Describe("Deputy", func() {
	Describe("Test diffing between Database and API", func() {
		Context("From nothing in DB", func() {
			It("should bring one diff that is an insert", func() {
				fromDB := []Models.Depute{}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(1))
				Expect(res[0].Operation).To(Equal("Create"))
				deputy := res[0].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})
			It("should bring one diff that is an insert with mutiple fields", func() {
				fromDB := []Models.Depute{}
				fromAPI := []Models.Depute{
					{
						Slug:    "lel",
						Twitter: "@lel",
						Activites: []Models.Activite{
							{
								PresencesCommission:      1,
								PresencesHemicycle:       2,
								ParticipationsCommission: 3,
								ParticipationsHemicycle:  4,
								Questions:                5,
								Vacances:                 6,
							},
						},
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(1))
			})
		})

		Context("With already created deputies in DB", func() {
			It("Should bring one diff that is an update", func() {
				fromDB := []Models.Depute{
					{
						Slug:        "lel",
						GroupeSigle: "LREM",
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug:        "lel",
						GroupeSigle: "FI",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(1))
				Expect(res[0].Operation).To(Equal("Update"))
				deputy := res[0].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
				Expect(deputy.GroupeSigle).To(Equal("FI"))
			})
		})

		Context("Deleting deputies in DB", func() {
			It("Should bring one diff that is an update", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				fromAPI := []Models.Depute{}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(1))
				Expect(res[0].Operation).To(Equal("Delete"))
				deputy := res[0].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})
		})
	})
})
