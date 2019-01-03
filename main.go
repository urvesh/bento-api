package main

import (
	"bento-api/config"
	"bento-api/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

func main() {
	config.InitMongo()

	//result := Image{}
	//id := "ncim7431"
	//err = collection.Find(bson.M{"_id": id}).One(&result)
	//
	//if err != nil {
	//	fmt.Println("error finding", err)
	//}
	//
	//fmt.Println("image", result)
	//
	//var docs []bson.M
	//
	//err = collection.Find(bson.M{"_id": id}).All(&docs)
	//
	//if err != nil {
	//	fmt.Println("find err", err)
	//}
	//
	//jsonResp, merr := json.Marshal(docs)
	//
	//if merr != nil {
	//	fmt.Println("marshal error", merr)
	//}
	//
	//fmt.Println("json:", string(jsonResp))

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: types.QueryType,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
		ResultCallbackFn: func(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte) {
			// TODO: NEED TO FIGURE OUT EXTENSIONS!!

			//type GResult struct {
			//	Data       interface{}                `json:"data"`
			//	Errors     []gqlerrors.FormattedError `json:"errors,omitempty"`
			//	Extensions interface{}                `json:"extensions"`
			//}
			//fmt.Println("result", string(responseBody))
			//
			//var a GResult
			//
			//a.Data = result.Data
			//a.Errors = result.Errors
			//a.Extensions = "ok"

		},
	})

	http.Handle("/", h)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		u4, _ := uuid.NewV4()

		result := executeQuery(r.URL.Query().Get("query"), schema)

		data := &GResult{
			Data: result.Data,
			Extensions: Extensions{
				CorrelationID: u4,
				RunTime:       time.Since(start) / 1000 / 1000,
			},
		}

		jsonResp, _ := json.Marshal(result)

		data.Extensions.Size = len(jsonResp)

		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("server starting on localhost:12345")

	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		fmt.Println("server error", err)
	}
}

// implementation using graphql-go, without graphiql, but includes extensions obj
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	data := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(data.Errors) > 0 {
		fmt.Println("Wrong result ", data.Errors)
	}

	return data
}

type Extensions struct {
	CorrelationID uuid.UUID     `json:"correlationID"`
	RunTime       time.Duration `json:"runTime"`
	Size          int           `json:"size"`
}

type GResult struct {
	Data       interface{}                 `json:"data,omitempty"`
	Errors     []gqlerrors.FormattedErrors `json:"errors,omitempty"`
	Extensions Extensions                  `json:"extensions"`
}
