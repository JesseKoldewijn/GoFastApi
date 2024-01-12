/*
Entrypoint for the GoFastApi application
*/
package main

import (
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/JesseKoldewijn/GoFastApi/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

var maxThreadsInt = runtime.NumCPU();
var maxThreadsString = strconv.Itoa(maxThreadsInt); // this returns an empty string for some reason

func main() {
	runtime.GOMAXPROCS(maxThreadsInt);
	
	println("Running on " + runtime.GOOS + " with " + maxThreadsString + " threads")
	
	App := fiber.New(fiber.Config{
		StreamRequestBody: true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName: "GoFastApi",
		
		JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
	})

	App.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	// Routes
	routes.Root(App) // /
	routes.Echo(App) // /echo/:message**
	
	log.Fatal(App.Listen(":8080"))
}