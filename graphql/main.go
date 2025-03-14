package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var config AppConfig
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}
	// Debugging: Print out config values
	log.Println("Account URL:", config.AccountURL)
	log.Println("Catalog URL:", config.CatalogURL)
	log.Println("Order URL:", config.OrderURL)

	s, err := NewGraphqlServer(config.AccountURL, config.CatalogURL, config.OrderURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.NewDefaultServer(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL", "/graphql"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
