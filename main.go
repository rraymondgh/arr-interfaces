package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/rraymondgh/arr-interfaces/internal/app"
)

func main() {
	app.New().Run()
}
