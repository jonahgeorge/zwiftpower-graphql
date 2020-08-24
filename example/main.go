package main

import (
	"log"
	"net/http"
	"os"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jonahgeorge/zwiftpower-go"
	"github.com/rs/cors"
)

type ZwiftPowerProxy struct {
	zpClient *zwiftpower.Client
}

var zpClient = &zwiftpower.Client{
	HTTPClient: &zwiftpower.LoggingHTTPClient{
		HTTPClient: http.DefaultClient,
	},
}

var zpp = &ZwiftPowerProxy{zpClient: zpClient}

func (zpp ZwiftPowerProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	schemaConfig := graphql.SchemaConfig{
		Query: RootQuery,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/graphiql", graphiqlHandler)
	mux.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))

	cors.Default().Handler(mux).ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on " + port)
	http.ListenAndServe(":"+port, zpp)
}
