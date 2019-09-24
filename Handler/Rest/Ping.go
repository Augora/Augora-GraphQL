package Rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")

	fmt.Fprintf(w, "pong")
}
