package main

import (
	"fmt"
	"log"

	"github.com/FatihKaanAkkus/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	err = cfg.SetUser("fatih")
	if err != nil {
		log.Fatalf("error setting user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)

	}

	fmt.Println(cfg)
}
