package types

import "github.com/graphql-go/graphql"

var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Description: "",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})
