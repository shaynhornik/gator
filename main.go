package main

import (
	"github.com/shaynhornik/gator/internal/config"
	"fmt"
	"log"
)

func main () {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("shayn")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
