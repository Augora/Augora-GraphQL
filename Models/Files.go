package Models

import "github.com/jinzhu/gorm"

type Dossiers struct {
	Sections []DossierHandler `json:"sections"`
}

type DossierHandler struct {
	Dossier Dossier `json:"section"`
}

type Dossier struct {
	gorm.Model
	IDFromAPI            uint   `json:"id"`
	IDDossierInstitution string `json:"id_dossier_institution"`
	Titre                string `json:"titre"`
	DateMinimum          string `json:"min_date"`
	DateMaximum          string `json:"max_date"`
	NombreIntervention   uint   `json:"nb_interventions"`
	URLInsitution        string `json:"url_institution"`
	URLNosDeputes        string `json:"url_nosdeputes"`
	URLNosDeputesAPI     string `json:"url_nosdeputes_api"`
}
