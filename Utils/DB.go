package Utils

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetDataBaseConnection() *gorm.DB {
	user := os.Getenv("BACKEND_SQL_USER")
	pass := os.Getenv("BACKEND_SQL_PASSWORD")
	database := os.Getenv("BACKEND_SQL_DATABASE")
	db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@augora-server.database.windows.net:1433?database="+database)
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)

	return db
}

func GetFakeLocalDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)

	return db
}
