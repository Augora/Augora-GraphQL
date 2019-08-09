package Models

import "github.com/jinzhu/gorm"

type Deputes struct {
	Deputes []DeputeHandler `json:"deputes"`
}

type Site struct {
	gorm.Model
	SiteRefer uint
	Site      string `json:"site"`
}

type Email struct {
	gorm.Model
	EmailRefer uint
	Email      string `json:"email"`
}

type Adresse struct {
	gorm.Model
	AdresseRefer uint
	Adresse      string `json:"adresse"`
}

type Collaborateur struct {
	gorm.Model
	CollaborateurRefer uint
	Collaborateur      string `json:"collaborateur"`
}

type DeputeHandler struct {
	Depute Depute `json:"depute"`
}

type Depute struct {
	gorm.Model
	IDFromAPI          uint            `json:"id"`
	Nom                string          `json:"nom"`
	NomDeFamille       string          `json:"nom_de_famille"`
	Prenom             string          `json:"prenom"`
	Sexe               string          `json:"sexe"`
	DateNaissance      string          `json:"date_naissance"`
	LieuNaissance      string          `json:"lieu_naissance"`
	NumDepartement     string          `json:"num_deptmt"`
	NomCirco           string          `json:"nom_circo"`
	NumCirco           int             `json:"num_circo"`
	MandatDebut        string          `json:"mandat_debut"`
	GroupeSigle        string          `json:"groupe_sigle"`
	PartiRattFinancier string          `json:"parti_ratt_financier"`
	SitesWeb           []Site          `json:"sites_web" gorm:"foreignkey:SiteRefer"`
	Emails             []Email         `json:"emails" gorm:"foreignkey:EmailRefer"`
	Adresses           []Adresse       `json:"adresses" gorm:"foreignkey:AdresseRefer"`
	Collaborateurs     []Collaborateur `json:"collaborateurs" gorm:"foreignkey:CollaborateurRefer"`
	Profession         string          `json:"profession"`
	PlaceEnHemicyle    string          `json:"place_en_hemicycle"`
	UrlAN              string          `json:"url_an"`
	IDAN               string          `json:"id_an"`
	Slug               string          `json:"slug"`
	UrlNosDeputes      string          `json:"url_nosdeputes"`
	UrlNosDeputesAPI   string          `json:"url_nosdeputes_api"`
	NombreMandats      int             `json:"nb_mandats"`
	Twitter            string          `json:"twitter"`
	EstEnMandat        bool
}
