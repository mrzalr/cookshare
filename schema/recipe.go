package schema

import "github.com/graphql-go/graphql"

var recipeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Recipe",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"userId": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"imageUrl": &graphql.Field{
				Type: graphql.String,
			},
			"ingredients": &graphql.Field{
				Type: graphql.String,
			},
			"instructions": &graphql.Field{
				Type: graphql.String,
			},
			"servings": &graphql.Field{
				Type: graphql.String,
			},
			"notes": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
