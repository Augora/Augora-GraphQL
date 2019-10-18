package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Deputes struct {
	Deputes []DeputeHandler `json:"deputes"`
}

type Site struct {
	gorm.Model `json:"-"`
	SiteRefer  uint   `json:"-"`
	Site       string `json:"site"`
}

type Email struct {
	gorm.Model `json:"-"`
	EmailRefer uint   `json:"-"`
	Email      string `json:"email"`
}

type Adresse struct {
	gorm.Model   `json:"-"`
	AdresseRefer uint   `json:"-"`
	Adresse      string `json:"adresse"`
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
	gorm.Model         `json:"-"`
	CollaborateurRefer uint   `json:"-"`
	Collaborateur      string `json:"collaborateur"`
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
	gorm.Model               `json:"-"`
	ActiviteRefer            uint      `json:"-"`
	DateDebut                time.Time `json:"date_debut"`
	DateFin                  time.Time `json:"date_fin"`
	NumeroDeSemaine          uint      `json:"numero_de_semaine"`
	PresencesCommission      uint      `json:"presences_commission"`
	PresencesHemicycle       uint      `json:"presences_hemicycle"`
	ParticipationsCommission uint      `json:"participations_commission"`
	ParticipationsHemicycle  uint      `json:"participations_hemicycle"`
	Questions                uint      `json:"questions"`
	Vacances                 uint      `json:"vacances"`
}

type Depute struct {
	gorm.Model `json:"-"`

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
	Slug               string `json:"slug"`
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
	Activites      []Activite      `gorm:"foreignkey:ActiviteRefer" json:"-"`

	// Custom fields
	EstEnMandat bool `json:"-"`
}
