package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/isso0424/portfolio_api/graph"
)

var schema = graph.CreateSchema()

func graphQLRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	authrozers := r.Header.Values("Authorize")
	if len(authrozers) == 0 {
		authrozers = append(authrozers, "")
	}

	token := authrozers[0]

	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: string(data),
			Context: context.WithValue(context.Background(), "token", token),
		},
	)
	if len(result.Errors) > 0 {
		log.Printf("failed to execute graphql operation.\nerror: %+v", result.Errors)
	}

	rJson, _ := json.Marshal(result)
	w.Write(rJson)
}

func Serve(port string) {
	http.HandleFunc("/graphql", graphQLRoute)
	http.ListenAndServe(port, nil)
}
