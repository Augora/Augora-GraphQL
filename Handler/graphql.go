package Handler

import (
	"context"
	"net/http"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type server struct {
}

func (s *server) registerQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()

	obj.FieldFunc("Deputes", func() []Models.Depute {
		var deputes []Models.Depute
		db := Utils.GetDataBaseConnection()
		defer db.Close()
		db.Set("gorm:auto_preload", true).Find(&deputes)
		return deputes
	})

	obj.FieldFunc("DeputesEnMandat", func() []Models.Depute {
		var deputes []Models.Depute
		db := Utils.GetDataBaseConnection()
		defer db.Close()
		db.Set("gorm:auto_preload", true).Where(&Models.Depute{EstEnMandat: true}).Find(&deputes)
		return deputes
	})

	obj.FieldFunc("Depute", func(args struct{ Slug string }) Models.Depute {
		var depute Models.Depute
		db := Utils.GetDataBaseConnection()
		defer db.Close()
		db.Set("gorm:auto_preload", true).Where(&Models.Depute{Slug: args.Slug}).Find(&depute)
		return depute
	})
}

func (s *server) registerMutation(schema *schemabuilder.Schema) {
	obj := schema.Mutation()
	obj.FieldFunc("ping", func(args struct{ Message string }) string {
		return args.Message
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
		for _, site := range m.Sites {
			result = append(result, site.Site)
		}
		return result, nil
	})

	object.FieldFunc("autresMandats", func(ctx context.Context, m *Models.Depute) ([]string, error) {
		var result []string
		for _, autreMandat := range m.AutresMandats {
			result = append(result, autreMandat.AutreMandat)
		}
		return result, nil
	})

	object.FieldFunc("anciensMandats", func(ctx context.Context, m *Models.Depute) ([]string, error) {
		var result []string
		for _, ancienMandat := range m.AnciensMandats {
			result = append(result, ancienMandat.AncienMandat)
		}
		return result, nil
	})
}

// schema builds the graphql schema.
func (s *server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.registerDepute(builder)
	s.registerQuery(builder)
	s.registerMutation(builder)
	return builder.MustBuild()
}

func GraphQLHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Instantiate a server, build a server, and serve the schema on port 3030.
	server := &server{}

	// Init schemas
	schema := server.schema()
	introspection.AddIntrospectionToSchema(schema)

	// Expose schema and graphiql.
	hdl := graphql.HTTPHandler(schema)
	hdl.ServeHTTP(w, r)
}
