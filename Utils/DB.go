package Utils

import (
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
