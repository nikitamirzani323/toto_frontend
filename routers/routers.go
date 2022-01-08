package routers

import (
	"bitbucket.org/isbtotogroup/frontend_svelte/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(etag.New())

	app.Static("/", "svelte/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controller.HealthCheck)
	app.Post("api/init", controller.InitToken)
	app.Post("api/listpasaran", controller.Listpasaran)
	app.Post("api/checkpasaran", controller.Checkpasaran)
	app.Post("api/inittogel_432d", controller.Inittogel_432d)
	app.Post("api/limittogel", controller.Clientlimitpasaran)
	app.Post("api/resulttogel", controller.Resulttogel)
	app.Post("api/resulttogelall", controller.ResulttogelAll)
	app.Post("api/invoicebet", controller.Invoicebet)
	app.Post("api/invoicebetdetail", controller.Invoicebetid)
	app.Post("api/slipperiode", controller.Slipperiode)
	app.Post("api/slipperiodeall", controller.SlipperiodeAll)
	app.Post("api/slipperiodedetail", controller.Slipperiodedetail)
	app.Post("api/savetransaksi", controller.Savetransaksi)
	app.Post("api/bookdream", controller.Listbukumimpi)
	return app
}
