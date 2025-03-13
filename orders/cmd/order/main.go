package main

import (
	"Goservices/orders"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:DATABASE_URL`
	AccountURL  string `envconfig:ACCOUNT_SERVICE	_URL`
	CatalogURL  string `envconfig:CATALOG_SERVICE_URL`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	var r orders.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error){
		r,err=orders.NewPostgresRepository(cfg.DatabaseURL)
		if err!=nil{
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8000")
	s :=orders.NewService(r)
	log.Fatal(orders.ListenGRPC(s,cfg,cfg.AccountURL,cfg,cfg.CatalogURL,8000))

}
