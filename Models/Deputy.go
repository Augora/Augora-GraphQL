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
	Site       string `json:"site"`
}

type Email struct {
	gorm.Model `json:"-" diff:"-"`
	EmailRefer uint   `json:"-" diff:"-"`
	Email      string `json:"email"`
}

type Adresse struct {
	gorm.Model   `json:"-" diff:"-"`
	AdresseRefer uint   `json:"-" diff:"-"`
	Adresse      string `json:"adresse"`
}

type Collaborateur struct {
	gorm.Model         `json:"-" diff:"-"`
	CollaborateurRefer uint   `json:"-" diff:"-"`
	Collaborateur      string `json:"collaborateur"`
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
	Slug               string `json:"slug"`
	UrlNosDeputes      string `json:"url_nosdeputes"`
	UrlNosDeputesAPI   string `json:"url_nosdeputes_api"`
	NombreMandats      int    `json:"nb_mandats"`
	Twitter            string `json:"twitter"`

	// ForeignKey fields
	SitesWeb       []Site          `json:"sites_web" gorm:"foreignkey:SiteRefer" diff:"-"`
	Emails         []Email         `json:"emails" gorm:"foreignkey:EmailRefer" diff:"-"`
	Adresses       []Adresse       `json:"adresses" gorm:"foreignkey:AdresseRefer" diff:"-"`
	Collaborateurs []Collaborateur `json:"collaborateurs" gorm:"foreignkey:CollaborateurRefer" diff:"-"`
	Activites      []Activity      `gorm:"foreignkey:ActivityRefer" json:"-" diff:"-"`

	// Custom fields
	EstEnMandat bool `json:"-"`
}
