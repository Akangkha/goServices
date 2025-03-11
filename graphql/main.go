package main

import (
	"log"
	"net/http"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
	"github.com/99designs/gqlgen/graphql/playground"
	
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
  
   s,err := NewGraphqlServer(config.AccountURL,config.CatalogURL,config.OrderURL)
   if err != nil {
	log.Fatal(err)
   }

   http.Handle("/graphql",handler.GraphQL(s.ExecutableSchema()))
   http.Handle("/playground",playground.Handler("GraphQL", "/graphql"))
   log.Fatal(http.ListenAndServe(":8080",nil))
}