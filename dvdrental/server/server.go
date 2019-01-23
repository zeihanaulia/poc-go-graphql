package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/lib/pq"
	"github.com/zeihanaulia/poc-go-graphql/dvdrental"
)

const defaultPort = "8080"

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gqlHandler := handler.GraphQL(
		dvdrental.NewExecutableSchema(dvdrental.Config{Resolvers: &dvdrental.Resolver{}}),
	)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", dvdrental.DataLoaderMiddleware(db, gqlHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
