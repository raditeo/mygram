package main

import (
	"fmt"
	"log"
	"mygram/app"
	"mygram/configs"
	"mygram/middlewares/notfoundhandler"
	"mygram/middlewares/recover"
	"mygram/routers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//prepare config
	configs.LoadConfig()

	//prepare fiber
	fiberApp := fiber.New(configs.FiberConfigFromEnv())

	fiberApp.Use(recover.New(configs.RecoverConfig()))

	//add routers
	routers.RegisterRouters(fiberApp)

	//register route not found
	fiberApp.Use(notfoundhandler.New())
	//gracefull shutdown
	go func() {
		if err := fiberApp.Listen(app.Port.String()); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // When an interru
	fmt.Println("Gracefully shutting down...")
	_ = fiberApp.Shutdown()
	fmt.Println("Running cleanup tasks...")
	fmt.Println("Fiber was successful shutdown.")
}
