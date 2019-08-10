package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

// server is our graphql server.
type server struct {
	db *gorm.DB
}

func GetDataBaseConnection() *gorm.DB {
	user := os.Getenv("backend_sql_user")
	pass := os.Getenv("backend_sql_password")
	db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@augora.database.windows.net:1433?database=augora-db")
	if err != nil {
		fmt.Println(err)
	}
	// db.LogMode(true)

	return db
}

// registerQuery registers the root query type.
func (s *server) registerQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()

	obj.FieldFunc("Deputes", func() []Models.Depute {
		var deputes []Models.Depute
		s.db.Set("gorm:auto_preload", true).Find(&deputes)
		return deputes
	})

	obj.FieldFunc("DeputesEnMandat", func(ctx context.Context) []Models.Depute {
		var deputes []Models.Depute
		s.db.Set("gorm:auto_preload", true).Where(&Models.Depute{EstEnMandat: true}).Find(&deputes)
		return deputes
	})

	obj.FieldFunc("Depute", func(args struct{ Slug string }) Models.Depute {
		var depute Models.Depute
		s.db.Set("gorm:auto_preload", true).Where(&Models.Depute{Slug: args.Slug}).Find(&depute)
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

	// Init schemas
	schema := server.schema()
	introspection.AddIntrospectionToSchema(schema)

	// Init Database connection
	server.db = GetDataBaseConnection()

	// Expose schema and graphiql.
	hdl := graphql.HTTPHandler(schema)
	hdl.ServeHTTP(w, r)
}
