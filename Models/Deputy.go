package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Deputes struct {
	Deputes []DeputeHandler `json:"deputes"`
}

type Site struct {
	gorm.Model `json:"-" diff:"-"`
	SiteRefer  uint   `json:"-" diff:"-"`
	Site       string `json:"site" diff:"Site,identifier"`
}

type Email struct {
	gorm.Model `json:"-" diff:"-"`
	EmailRefer uint   `json:"-" diff:"-"`
	Email      string `json:"email" diff:"Email,identifier"`
}

type Adresse struct {
	gorm.Model   `json:"-" diff:"-"`
	AdresseRefer uint   `json:"-" diff:"-"`
	Adresse      string `json:"adresse" diff:"Adresse,identifier"`
}

type AncienMandat struct {
	gorm.Model        `json:"-" diff:"-"`
	AncienMandatRefer uint      `json:"-" diff:"-"`
	AncienMandat      string    `json:"mandat" diff:"AncienMandat,identifier"`
	DateDebut         time.Time `diff:"-"`
	DateFin           time.Time `diff:"-"`
	Intitule          string    `diff:"-"`
}

type AutreMandat struct {
	gorm.Model       `json:"-" diff:"-"`
	AutreMandatRefer uint      `json:"-" diff:"-"`
	AutreMandat      string    `json:"mandat" diff:"AutreMandat,identifier"`
	DateDebut        time.Time `diff:"-"`
	DateFin          time.Time `diff:"-"`
	Intitule         string    `diff:"-"`
}

type Collaborateur struct {
	gorm.Model         `json:"-" diff:"-"`
	CollaborateurRefer uint   `json:"-" diff:"-"`
	Collaborateur      string `json:"collaborateur" diff:"Collaborateur,identifier"`
}

type DeputeHandler struct {
	Depute Depute `json:"depute"`
}

type ActivitesHandler struct {
	DateDebut     string     `json:"date_debut"`
	DateDebutParl string     `json:"date_debut_parl"`
	DateFin       string     `json:"date_fin"`
	Data          []Activite `json:"data"`
}

type Activite struct {
	gorm.Model               `json:"-" diff:"-"`
	ActiviteRefer            uint      `diff:"-"`
	DateDebut                time.Time `json:"date_debut" diff:"-"`
	DateFin                  time.Time `json:"date_fin" diff:"-"`
	NumeroDeSemaine          uint      `json:"numero_de_semaine" diff:"NumeroDeSemaine,identifier"`
	PresencesCommission      uint      `json:"presences_commission"`
	PresencesHemicycle       uint      `json:"presences_hemicycle"`
	ParticipationsCommission uint      `json:"participations_commission"`
	ParticipationsHemicycle  uint      `json:"participations_hemicycle"`
	Questions                uint      `json:"questions"`
	Vacances                 uint      `json:"vacances"`
}

type Depute struct {
	gorm.Model `diff:"-"`

	// Fields from API
	Nom                string `json:"nom"`
	NomDeFamille       string `json:"nom_de_famille"`
	Prenom             string `json:"prenom"`
	Sexe               string `json:"sexe"`
	DateNaissance      string `json:"date_naissance"`
	LieuNaissance      string `json:"lieu_naissance"`
	NumDepartement     string `json:"num_deptmt"`
	NomCirco           string `json:"nom_circo"`
	NumCirco           int    `json:"num_circo"`
	MandatDebut        string `json:"mandat_debut"`
	GroupeSigle        string `json:"groupe_sigle"`
	PartiRattFinancier string `json:"parti_ratt_financier"`
	Profession         string `json:"profession"`
	PlaceEnHemicyle    string `json:"place_en_hemicycle"`
	UrlAN              string `json:"url_an"`
	IDAN               string `json:"id_an"`
	Slug               string `json:"slug" diff:"Slug,identifier"`
	UrlNosDeputes      string `json:"url_nosdeputes"`
	UrlNosDeputesAPI   string `json:"url_nosdeputes_api"`
	NombreMandats      int    `json:"nb_mandats"`
	Twitter            string `json:"twitter"`

	// ForeignKey fields
	Sites          []Site          `json:"sites_web" gorm:"foreignkey:SiteRefer"`
	Emails         []Email         `json:"emails" gorm:"foreignkey:EmailRefer"`
	Adresses       []Adresse       `json:"adresses" gorm:"foreignkey:AdresseRefer"`
	Collaborateurs []Collaborateur `json:"collaborateurs" gorm:"foreignkey:CollaborateurRefer"`
	AnciensMandats []AncienMandat  `json:"anciens_mandats" gorm:"foreignkey:AncienMandatRefer"`
	AutresMandats  []AutreMandat   `json:"autres_mandats" gorm:"foreignkey:AutreMandatRefer"`
	Activites      []Activite      `json:"-" gorm:"foreignkey:ActiviteRefer"`

	// Custom fields
	EstEnMandat bool `json:"-"`
}

type GenericDiff struct {
	Operation string
	Item      interface{}
}

func MergeDeputies(deputyFromDB Depute, deputyFromAPI Depute) Depute {
	var newDeputy Depute
	newDeputy = deputyFromAPI
	newDeputy.ID = deputyFromDB.ID

	// SitesWeb
	for newSiteIdx := range newDeputy.Sites {
		newDeputy.Sites[newSiteIdx].ID = 0
		for dbSiteIdx := range deputyFromDB.Sites {
			if newDeputy.Sites[newSiteIdx].Site == deputyFromDB.Sites[dbSiteIdx].Site {
				newDeputy.Sites[newSiteIdx].ID = deputyFromDB.Sites[dbSiteIdx].ID
			}
		}
	}

	// Emails
	for newEmailIdx := range newDeputy.Emails {
		newDeputy.Emails[newEmailIdx].ID = 0
		for dbEmailIdx := range deputyFromDB.Emails {
			if newDeputy.Emails[newEmailIdx].Email == deputyFromDB.Emails[dbEmailIdx].Email {
				newDeputy.Emails[newEmailIdx].ID = deputyFromDB.Emails[dbEmailIdx].ID
			}
		}
	}

	// Adresses
	for newAdresseIdx := range newDeputy.Adresses {
		newDeputy.Adresses[newAdresseIdx].ID = 0
		for dbAdresseIdx := range deputyFromDB.Adresses {
			if newDeputy.Adresses[newAdresseIdx].Adresse == deputyFromDB.Adresses[dbAdresseIdx].Adresse {
				newDeputy.Adresses[newAdresseIdx].ID = deputyFromDB.Adresses[dbAdresseIdx].ID
			}
		}
	}

	// Collaborateurs
	for newCollaborateurIdx := range newDeputy.Collaborateurs {
		newDeputy.Collaborateurs[newCollaborateurIdx].ID = 0
		for dbCollaborateurIdx := range deputyFromDB.Collaborateurs {
			if newDeputy.Collaborateurs[newCollaborateurIdx].Collaborateur == deputyFromDB.Collaborateurs[dbCollaborateurIdx].Collaborateur {
				newDeputy.Collaborateurs[newCollaborateurIdx].ID = deputyFromDB.Collaborateurs[dbCollaborateurIdx].ID
			}
		}
	}

	// AnciensMandats
	for newAnciensMandatIdx := range newDeputy.AnciensMandats {
		newDeputy.AnciensMandats[newAnciensMandatIdx].ID = 0
		for dbAnciensMandatIdx := range deputyFromDB.AnciensMandats {
			if newDeputy.AnciensMandats[newAnciensMandatIdx].AncienMandat == deputyFromDB.AnciensMandats[dbAnciensMandatIdx].AncienMandat {
				newDeputy.AnciensMandats[newAnciensMandatIdx].ID = deputyFromDB.AnciensMandats[dbAnciensMandatIdx].ID
			}
		}
	}

	// AutresMandats
	for newAutresMandatIdx := range newDeputy.AutresMandats {
		newDeputy.AutresMandats[newAutresMandatIdx].ID = 0
		for dbAutresMandatIdx := range deputyFromDB.AutresMandats {
			if newDeputy.AutresMandats[newAutresMandatIdx].AutreMandat == deputyFromDB.AutresMandats[dbAutresMandatIdx].AutreMandat {
				newDeputy.AutresMandats[newAutresMandatIdx].ID = deputyFromDB.AutresMandats[dbAutresMandatIdx].ID
			}
		}
	}

	// Activites
	for newActivitesIdx := range newDeputy.Activites {
		newDeputy.Activites[newActivitesIdx].ID = 0
		for dbActivitesIdx := range deputyFromDB.Activites {
			if newDeputy.Activites[newActivitesIdx].NumeroDeSemaine == deputyFromDB.Activites[dbActivitesIdx].NumeroDeSemaine {
				newDeputy.Activites[newActivitesIdx].ID = deputyFromDB.Activites[dbActivitesIdx].ID
			}
		}
	}

	return newDeputy
}
