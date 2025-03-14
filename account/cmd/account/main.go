package main

import (
	"Goservices/account"
	"log"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database URL: %s", cfg.DatabaseURL) // Log the database URL

	// Parse the URL to check for formatting issues
	parsedURL, err := url.Parse(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Invalid database URL format: %v", err)
	}
	log.Printf("Parsed URL: %v", parsedURL)

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("Failed to connect to the database:", err) // Log connection attempts
		} else {
			log.Println("Successfully connected to the database") // Log successful connection
		}
		return
	})
	defer r.Close()
	log.Println("Listening to port 8000......")
	s := account.NewService(r)
	log.Print(s)
	log.Fatal(account.ListenGRPC(s, 8000))
}
