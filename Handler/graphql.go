package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/KevinBacas/Gin-Go-Test/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type post struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

// server is our graphql server.
type server struct {
}

var hidden_db *gorm.DB
var isDBInitialized bool = false

func GetDataBaseConnection() *gorm.DB {
	if !isDBInitialized {
		user := os.Getenv("backend_sql_user")
		pass := os.Getenv("backend_sql_password")
		db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@augora.database.windows.net:1433?database=augora-db")
		if err != nil {
			fmt.Println(err)
		}
		db.LogMode(true)
		hidden_db = db
		isDBInitialized = true
	}

	return hidden_db
}

// registerQuery registers the root query type.
func (s *server) registerQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()

	obj.FieldFunc("Deputes", func() []Models.Depute {
		db := GetDataBaseConnection()

		var deputes []Models.Depute
		db.Set("gorm:auto_preload", true).Find(&deputes)
		return deputes
	})

	obj.FieldFunc("DeputesEnMandat", func(ctx context.Context) []Models.Depute {
		db := GetDataBaseConnection()

		var deputes []Models.Depute
		db.Set("gorm:auto_preload", true).Where(&Models.Depute{EstEnMandat: true}).Find(&deputes)
		return deputes
	})

	obj.FieldFunc("Depute", func(args struct{ Slug string }) Models.Depute {
		db := GetDataBaseConnection()

		var depute Models.Depute
		db.Set("gorm:auto_preload", true).Where(&Models.Depute{Slug: args.Slug}).Find(&depute)
		return depute
	})
}

func (s *server) registerDepute(schema *schemabuilder.Schema) {
	object := schema.Object("Depute", Models.Depute{})
	object.Description = "A single depute."

	object.FieldFunc("emails", func(ctx context.Context, m *Models.Depute) ([]string, error) {
		var result []string
		for _, email := range m.Emails {
			result = append(result, email.Email)
		}
		return result, nil
	})

	object.FieldFunc("collaborateurs", func(ctx context.Context, m *Models.Depute) ([]string, error) {
		var result []string
		for _, collaborateur := range m.Collaborateurs {
			result = append(result, collaborateur.Collaborateur)
		}
		return result, nil
	})

	object.FieldFunc("adresses", func(ctx context.Context, m *Models.Depute) ([]string, error) {
		var result []string
		for _, adresse := range m.Adresses {
			result = append(result, adresse.Adresse)
		}
		return result, nil
	})

	object.FieldFunc("sitesWeb", func(ctx context.Context, m *Models.Depute) ([]string, error) {
		var result []string
		for _, site := range m.SitesWeb {
			result = append(result, site.Site)
		}
		return result, nil
	})
}

// schema builds the graphql schema.
func (s *server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.registerDepute(builder)
	s.registerQuery(builder)
	return builder.MustBuild()
}

func GraphQLHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Instantiate a server, build a server, and serve the schema on port 3030.
	server := &server{}

	schema := server.schema()
	introspection.AddIntrospectionToSchema(schema)

	// Expose schema and graphiql.
	hdl := graphql.HTTPHandler(schema)
	hdl.ServeHTTP(w, r)
}
