package main

import (
	"log"
	"restaurant-order-system/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatalf("Error setting up app: %v", err)
	}
}
