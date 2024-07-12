package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/rraymondgh/arr-interface/internal/app"
)

func main() {
	app.New().Run()
}
