package types

import (
	"github.com/graphql-go/graphql"
)

var HeadlineType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Headline",
	Description: "",
	Fields: graphql.Fields{
		"primary": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				headline := p.Source.(string)
				return headline, nil
			},
		},
		"seo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				headline := p.Source.(string)
				return headline, nil
			},
		},
		"social": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				headline := p.Source.(string)
				return headline, nil
			},
		},
		"tease": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				headline := p.Source.(string)
				return headline, nil
			},
		},
	},
})
