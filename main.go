package main

import (
	"os"
	"user_list_go/bootstrap"
	"user_list_go/routes"
)

func init() {
	bootstrap.LoadEnv()
	bootstrap.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")
	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + port))
}
