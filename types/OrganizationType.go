package types

import "github.com/graphql-go/graphql"

var OrganizationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Organization",
	Description: "",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"copyright": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})
