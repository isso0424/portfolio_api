package main

import (
	"os"

	"github.com/isso0424/portfolio_api/auth"
	"github.com/isso0424/portfolio_api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	auth.SetSecret(os.Getenv("SECRET"))
	auth.SetKey(os.Getenv("KEY"))
	server.Serve(":8000")
}
