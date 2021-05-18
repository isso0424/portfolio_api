package main

import (
	"os"

	"github.com/isso0424/portfolio_api/auth"
	"github.com/isso0424/portfolio_api/db/sqlite"
	"github.com/isso0424/portfolio_api/graph/variables"
	"github.com/isso0424/portfolio_api/mem_db"
	"github.com/isso0424/portfolio_api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	switch (os.Getenv("DB")) {
	case "sqlite":
		dbSet, err := sqlite.Initialize(os.Getenv("DB_PATH"))
		if err != nil {
			panic(err)
		}
		variables.ContactDB = &dbSet.Contact
		variables.ProductDB = &dbSet.Product
	default:
		variables.ContactDB = &mem_db.ContactDB{}
		variables.ProductDB = &mem_db.ProductDB{}
		variables.SkillDB = &mem_db.SkillDB{}
	}

	auth.SetSecret(os.Getenv("SECRET"))
	auth.SetKey(os.Getenv("KEY"))
	server.Serve(":8000")
}
