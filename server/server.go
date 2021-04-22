package server

import (
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

	result := graphql.Do(
		graphql.Params{
			Schema: schema,
			RequestString: string(data),
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
