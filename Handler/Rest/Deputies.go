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

func GetDataBaseConnection() *gorm.DB {
	user := os.Getenv("backend_sql_user")
	pass := os.Getenv("backend_sql_password")
	db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@augora.database.windows.net:1433?database=augora-db")
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)

	return db
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600")
	db := GetDataBaseConnection()

	var deputes []Models.Depute
	db.Preload("sites").Preload("emails").Preload("adresses").Preload("collaborateurs").Find(&deputes)
	res, _ := json.Marshal(deputes)
	fmt.Fprintf(w, string(res))
}
