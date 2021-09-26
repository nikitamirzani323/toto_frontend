package main

import (
	"log"

	"github.com/nikitamirzani323/gofiber-frontendtogel/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
