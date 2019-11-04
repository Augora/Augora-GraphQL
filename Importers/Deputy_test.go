package Importers_test

import (
	"time"

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
				deputy := res[0].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
				Expect(deputy.Twitter).To(Equal("@lel"))
				Expect(deputy.Activites[0].PresencesCommission).To(Equal(uint(1)))
				Expect(deputy.Activites[0].PresencesHemicycle).To(Equal(uint(2)))
				Expect(deputy.Activites[0].ParticipationsCommission).To(Equal(uint(3)))
				Expect(deputy.Activites[0].ParticipationsHemicycle).To(Equal(uint(4)))
				Expect(deputy.Activites[0].Questions).To(Equal(uint(5)))
				Expect(deputy.Activites[0].Vacances).To(Equal(uint(6)))
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
			It("Should bring one diff that is an delete of a deputy", func() {
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

			It("Should bring one diff that is an delete of a Site", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						Sites: []Models.Site{
							{
								Site: "http://google.fr",
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				site := res[0].Item.(Models.Site)
				Expect(site.Site).To(Equal("http://google.fr"))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})

			It("Should bring one diff that is an delete of an Email", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						Emails: []Models.Email{
							{
								Email: "lel@mdr.eu",
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				email := res[0].Item.(Models.Email)
				Expect(email.Email).To(Equal("lel@mdr.eu"))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})

			It("Should bring one diff that is an delete of an Adresse", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						Adresses: []Models.Adresse{
							{
								Adresse: "8 rue keks",
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				adresse := res[0].Item.(Models.Adresse)
				Expect(adresse.Adresse).To(Equal("8 rue keks"))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})

			It("Should bring one diff that is an delete of a Collaborateur", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						Collaborateurs: []Models.Collaborateur{
							{
								Collaborateur: "Jean DUPONT",
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				collaborateur := res[0].Item.(Models.Collaborateur)
				Expect(collaborateur.Collaborateur).To(Equal("Jean DUPONT"))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})

			It("Should bring one diff that is an delete of an AncienMandat", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						AnciensMandats: []Models.AncienMandat{
							{
								AncienMandat: "10/10/1010 / ",
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				ancienMandat := res[0].Item.(Models.AncienMandat)
				Expect(ancienMandat.AncienMandat).To(Equal("10/10/1010 / "))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})

			It("Should bring one diff that is an delete of an AutreMandat", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						AutresMandats: []Models.AutreMandat{
							{
								AutreMandat: "10/10/1010 / ",
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				autreMandat := res[0].Item.(Models.AutreMandat)
				Expect(autreMandat.AutreMandat).To(Equal("10/10/1010 / "))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})

			It("Should bring one diff that is an delete of an AutreMandat", func() {
				fromDB := []Models.Depute{
					{
						Slug: "lel",
						Activites: []Models.Activite{
							{
								Vacances: 5,
							},
						},
					},
				}
				fromAPI := []Models.Depute{
					{
						Slug: "lel",
					},
				}
				res := DiffFromDB(fromDB, fromAPI)
				Expect(len(res)).To(Equal(2))
				Expect(res[0].Operation).To(Equal("Delete"))
				activite := res[0].Item.(Models.Activite)
				Expect(activite.Vacances).To(Equal(uint(5)))
				deputy := res[1].Item.(Models.Depute)
				Expect(deputy.Slug).To(Equal("lel"))
			})
		})
	})

	Describe("Test activites processing", func() {
		Context("Date processing", func() {
			It("Only one activity with an end date on a Monday", func() {
				activityHandler := Models.ActivitesHandler{
					DateDebut:     "2019-01-01",
					DateDebutParl: "2019-01-01",
					DateFin:       "2019-10-01",
					Data: []Models.Activite{
						{
							NumeroDeSemaine:          1,
							PresencesCommission:      2,
							PresencesHemicycle:       3,
							ParticipationsCommission: 4,
							ParticipationsHemicycle:  5,
							Questions:                6,
							Vacances:                 7,
						},
					},
				}
				processedActivities := ProccessActivitiesDates(activityHandler)
				Expect(processedActivities.Data[0].DateDebut).To(Equal(time.Date(2018, time.September, 24, 0, 0, 0, 0, time.UTC)))
				Expect(processedActivities.Data[0].DateFin).To(Equal(time.Date(2018, time.October, 1, 0, 0, 0, 0, time.UTC)))
			})
		})
	})

	Describe("Test mandat processing", func() {
		Context("AutreMandat", func() {
			It("One Mandat with all data", func() {
				mandat := Models.AutreMandat{
					AutreMandat: "Alfortville / Conseil municipal / membre",
				}
				newMandat := ProcessAutreMandat(mandat)
				Expect(newMandat.Localite).To(Equal("Alfortville"))
				Expect(newMandat.Institution).To(Equal("Conseil municipal"))
				Expect(newMandat.Intitule).To(Equal("membre"))
			})

			It("One Mandat with partial data", func() {
				mandat := Models.AutreMandat{
					AutreMandat: "Alfortville /  / membre",
				}
				newMandat := ProcessAutreMandat(mandat)
				Expect(newMandat.Localite).To(Equal("Alfortville"))
				Expect(newMandat.Institution).To(Equal(""))
				Expect(newMandat.Intitule).To(Equal("membre"))
			})

			It("Test the non-mutation of original object", func() {
				mandat := Models.AutreMandat{
					AutreMandat: "Alfortville / Conseil municipal / membre",
				}
				newMandat := ProcessAutreMandat(mandat)
				Expect(newMandat.Localite).To(Equal("Alfortville"))
				Expect(newMandat.Institution).To(Equal("Conseil municipal"))
				Expect(newMandat.Intitule).To(Equal("membre"))
				Expect(mandat.Localite).To(Equal(""))
				Expect(mandat.Institution).To(Equal(""))
				Expect(mandat.Intitule).To(Equal(""))
			})
		})

		Context("AncienMandat", func() {
			It("One Mandat with all data", func() {
				mandat := Models.AncienMandat{
					AncienMandat: "20/06/2007 / 19/06/2012 / fin de législature",
				}
				newMandat := ProcessAncienMandat(mandat)
				Expect(newMandat.DateDebut).To(Equal(time.Date(2007, time.June, 20, 0, 0, 0, 0, time.UTC)))
				Expect(newMandat.DateFin).To(Equal(time.Date(2012, time.June, 19, 0, 0, 0, 0, time.UTC)))
				Expect(newMandat.Intitule).To(Equal("fin de législature"))
			})

			It("One Mandat with partial data", func() {
				mandat := Models.AncienMandat{
					AncienMandat: "20/06/2007 /  / ",
				}
				newMandat := ProcessAncienMandat(mandat)
				Expect(newMandat.DateDebut).To(Equal(time.Date(2007, time.June, 20, 0, 0, 0, 0, time.UTC)))
				Expect(newMandat.DateFin).To(Equal(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)))
				Expect(newMandat.Intitule).To(Equal(""))
			})

			It("Test the non-mutation of original object", func() {
				mandat := Models.AncienMandat{
					AncienMandat: "20/06/2007 / 19/06/2012 / fin de législature",
				}
				newMandat := ProcessAncienMandat(mandat)
				Expect(newMandat.DateDebut).To(Equal(time.Date(2007, time.June, 20, 0, 0, 0, 0, time.UTC)))
				Expect(newMandat.DateFin).To(Equal(time.Date(2012, time.June, 19, 0, 0, 0, 0, time.UTC)))
				Expect(newMandat.Intitule).To(Equal("fin de législature"))
				Expect(mandat.DateDebut).To(Equal(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)))
				Expect(mandat.DateFin).To(Equal(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)))
				Expect(mandat.Intitule).To(Equal(""))
			})
		})
	})
})
