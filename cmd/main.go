package main

import (
	"hte-status-ms/internal/metrics"
	"hte-status-ms/internal/router"
	"log"
)

func main() {
	metrics.StartServer()

	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
