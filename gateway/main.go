package main

import (
	"gateway/internal/router"
	"log"
)

func main() {
	r := router.New()
	err := r.Run()
	if err != nil {
		log.Fatalf("failed to run server, %v", err)
	}
}