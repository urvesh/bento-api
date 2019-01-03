package types

import (
	"bento-api/config"
	"bento-api/lib"
	"bento-api/models"
	"fmt"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"image": &graphql.Field{
			Type: ImageType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, isOK := p.Args["id"].(string)

				fmt.Println("idQuery", id)

				if !isOK {
					fmt.Println("id was not passed in...")
					return nil, nil
				}

				query := config.Store.Image.Find(bson.M{"_id": id})

				var image models.Image
				var doc bson.M

				err := query.One(&image)
				_ = query.One(&doc)

				lib.PrintJSON("mongo image data", doc)

				if err != nil {
					fmt.Println("error from mongo", err)
					return nil, err
				}

				lib.PrintJSON("struct image data", image)

				return image, nil
			},
		},
	},
})
