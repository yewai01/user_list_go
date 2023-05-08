package main

import (
	"os"
	"user_list_go/bootstrap"
	"user_list_go/database"
	"user_list_go/routes"
)

func init() {
	bootstrap.LoadEnv()
	bootstrap.ConnectToDB()
	database.Migrate()
}

func main() {
	port := os.Getenv("PORT")
	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + port))
}
