package main

import (
	"Goservices/catalog"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cf Config
	err := envconfig.Process("", &cf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database URL: ", cf.DatabaseURL)

	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cf.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := catalog.NewService(r)

	log.Println("Service created in main.go ", s)
	log.Fatal(catalog.ListenGRPC(s, "8000"))
}
