package types

import (
	"bento-api/models"
	"github.com/graphql-go/graphql"
)

var ImageType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Image",
	Description: "A published news image",
	Fields: graphql.Fields{
		"altText": &graphql.Field{
			Type:        graphql.String,
			Description: "The alt text of the image, used for front-end presentation.",
		},
		"authors": &graphql.Field{
			Type:        graphql.NewList(PersonType),
			Description: "The Person resources associated as Authors.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resource := p.Source.(models.Image)
				return resource.Authors, nil
			},
		},
		"caption": &graphql.Field{
			Type: graphql.String,
		},
		"dateCreated": &graphql.Field{
			Type: graphql.String,
		},
		"dateModified": &graphql.Field{
			Type: graphql.String,
		},
		"datePublished": &graphql.Field{
			Type: graphql.String,
		},
		"encodingFormat": &graphql.Field{
			Type: graphql.String,
		},
		"graphicContent": &graphql.Field{
			Type: graphql.Boolean,
		},
		"headline": &graphql.Field{
			Type: HeadlineType,
		},
		"height": &graphql.Field{
			Type: graphql.Int,
		},
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"publisher": &graphql.Field{
			Type: OrganizationType,
		},
		"source": &graphql.Field{
			Type: OrganizationType,
		},
		"type": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "image", nil
			},
		},
		"url": &graphql.Field{
			Type: graphql.NewNonNull(UrlType),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				resource := p.Source.(models.Image)
				return resource, nil
			},
		},
		"width": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
