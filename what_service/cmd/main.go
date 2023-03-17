package main

import (
	_ "github.com/lib/pq"

	"user_service/internal/app"
)

const serviceName = "what_service"

func main() {
	a := app.New(serviceName)
	a.Run()
}
