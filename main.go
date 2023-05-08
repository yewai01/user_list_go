package main

import (
	"user_list_go/routes"
)

func main() {

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
