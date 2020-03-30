package main

import (
	"net/http"

	"github.com/heptiolabs/healthcheck"
)

func main() {
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	go http.ListenAndServe(":8080", health)
}
 