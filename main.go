package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/mrzalr/cookshare/schema"
)

func main() {
	//load dotenv files
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env files | err : %v", err)
	}

	//generating graphql schema
	schema, err := schema.GenerateSchema()
	if err != nil {
		log.Fatalf("Error load schema | err : %v", err)
	}

	//create graphql handler
	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	//register the handler into "/graphql" route
	http.Handle("/graphql", h)

	//start the server at localhost:$PORT
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")),
			nil,
		),
	)
}
