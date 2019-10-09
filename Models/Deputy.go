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
	gorm.Model        `json:"-"`
	AncienMandatRefer uint   `json:"-"`
	AncienMandat      string `json:"mandat"`
}

type AutreMandat struct {
	gorm.Model       `json:"-"`
	AutreMandatRefer uint   `json:"-"`
	AutreMandat      string `json:"mandat"`
}

type Collaborateur struct {
	gorm.Model         `json:"-" diff:"-"`
	CollaborateurRefer uint   `json:"-" diff:"-"`
	Collaborateur      string `json:"collaborateur" diff:"Collaborateur,identifier"`
}

type DeputeHandler struct {
	Depute Depute `json:"depute"`
}

type Activity struct {
	gorm.Model               `json:"-" diff:"-"`
	ActivityRefer            uint `json:"-" diff:"-"`
	StartDate                *time.Time
	EndDate                  *time.Time
	PresencesCommission      uint `json:"presencesCommission"`
	PresencesHemicycle       uint `json:"presencesHemicycle"`
	ParticipationsCommission uint `json:"participationsCommission"`
	ParticipationsHemicycle  uint `json:"participationsHemicycle"`
	Questions                uint `json:"questions"`
	Vacances                 uint `json:"vacances"`
}

type Depute struct {
	gorm.Model `json:"-" diff:"-"`

	// Fields from API
	IDFromAPI          uint   `json:"id"`
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
	SitesWeb       []Site          `json:"sites_web" gorm:"foreignkey:SiteRefer"`
	Emails         []Email         `json:"emails" gorm:"foreignkey:EmailRefer"`
	Adresses       []Adresse       `json:"adresses" gorm:"foreignkey:AdresseRefer"`
	Collaborateurs []Collaborateur `json:"collaborateurs" gorm:"foreignkey:CollaborateurRefer"`
	AnciensMandats []AncienMandat  `json:"anciens_mandats" gorm:"foreignkey:AncienMandatRefer"`
	AutresMandats  []AutreMandat   `json:"autres_mandats" gorm:"foreignkey:AutreMandatRefer"`
	Activites      []Activity      `gorm:"foreignkey:ActivityRefer" json:"-"`

	// Custom fields
	EstEnMandat bool `json:"-"`
}

type DeputyDiff struct {
	Operation string
	Deputy    Depute
}

func MergeDeputies(deputyFromDB Depute, deputyFromAPI Depute) Depute {
	var newDeputy Depute
	newDeputy = deputyFromAPI
	newDeputy.ID = deputyFromDB.ID

	// SitesWeb
	for newSiteIdx := range newDeputy.SitesWeb {
		for dbSiteIdx := range deputyFromDB.SitesWeb {
			if newDeputy.SitesWeb[newSiteIdx].Site == deputyFromDB.SitesWeb[dbSiteIdx].Site {
				newDeputy.SitesWeb[newSiteIdx].ID = deputyFromDB.SitesWeb[dbSiteIdx].ID
			}
		}
	}

	// Emails
	for newEmailIdx := range newDeputy.Emails {
		for dbEmailIdx := range deputyFromDB.Emails {
			if newDeputy.Emails[newEmailIdx].Email == deputyFromDB.Emails[dbEmailIdx].Email {
				newDeputy.Emails[newEmailIdx].ID = deputyFromDB.Emails[dbEmailIdx].ID
			}
		}
	}

	// Adresses
	for newAdresseIdx := range newDeputy.Adresses {
		for dbAdresseIdx := range deputyFromDB.Adresses {
			if newDeputy.Adresses[newAdresseIdx].Adresse == deputyFromDB.Adresses[dbAdresseIdx].Adresse {
				newDeputy.Adresses[newAdresseIdx].ID = deputyFromDB.Adresses[dbAdresseIdx].ID
			}
		}
	}

	// Collaborateurs
	for newCollaborateurIdx := range newDeputy.Collaborateurs {
		for dbCollaborateurIdx := range deputyFromDB.Collaborateurs {
			if newDeputy.Collaborateurs[newCollaborateurIdx].Collaborateur == deputyFromDB.Collaborateurs[dbCollaborateurIdx].Collaborateur {
				newDeputy.Collaborateurs[newCollaborateurIdx].ID = deputyFromDB.Collaborateurs[dbCollaborateurIdx].ID
			}
		}
	}

	return newDeputy
}
