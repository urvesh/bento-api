package types

import (
	"bento-api/models"
	"github.com/graphql-go/graphql"
	"strings"
)

var UrlType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Url",
	Description: "",
	Fields: graphql.Fields{
		"canonical": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resource := p.Source.(models.Image)
				return resource.Url, nil
			},
		},
		"primary": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resource := p.Source.(models.Image)
				return resource.Url, nil
			},
		},
		"short": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resource := p.Source.(models.Image)
				return resource.Url, nil
			},
		},
		"slug": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resource := p.Source.(models.Image)
				stringArray := strings.Split(resource.Url, "/")
				return stringArray[len(stringArray)-1], nil
			},
		},
	},
})
