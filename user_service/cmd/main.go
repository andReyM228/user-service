package main

import (
	"user_service/internal/app"

	_ "github.com/lib/pq"
)

const serviceName = "user_service"

func main() {
	a := app.New(serviceName)
	a.Run()
}
