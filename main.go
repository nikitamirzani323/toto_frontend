package main

import (
	"log"

	"bitbucket.org/isbtotogroup/frontend_svelte/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
