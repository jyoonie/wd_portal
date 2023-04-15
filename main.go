package main

import (
	"log"
	"portal/service"
)

func main() {
	s, err := service.New()
	if err != nil {
		log.Fatal("error creating service: %w", err)
	}

	if s.ListenAndServe("0.0.0.0:8080") != nil {
		log.Fatal("error listening and serving: %w", err)
	}
}
