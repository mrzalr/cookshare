package schema

import "github.com/graphql-go/graphql"

type graphqlField struct {
	Output      graphql.Output
	ResolveFn   graphql.FieldResolveFn
	Description string
	Args        graphql.FieldConfigArgument
}

func (gField *graphqlField) GenerateField() *graphql.Field {
	return &graphql.Field{
		Type:        gField.Output,
		Args:        gField.Args,
		Resolve:     gField.ResolveFn,
		Description: gField.Description,
	}
}

func GenerateSchema() (*graphql.Schema, error) {
	userType := graphqlField{
		Output: userType,
		ResolveFn: func(p graphql.ResolveParams) (interface{}, error) {
			return struct {
				Id        string `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				CreatedAt string `json:"createdAt"`
			}{"7d23-j81b-3ss6", "test", "test@test.com", "1980-01-01T13.00.00"}, nil
		},
		Description: "List of users",
		Args:        nil,
	}

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"user": userType.GenerateField(),
			},
		},
	)

	// rootMutation := graphql.NewObject(
	// 	graphql.ObjectConfig{
	// 		Name:   "RootMutation",
	// 		Fields: graphql.Fields{},
	// 	},
	// )

	schemaConfig := graphql.SchemaConfig{
		Query:        rootQuery,
		Mutation:     nil,
		Subscription: nil,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	return &schema, err
}
