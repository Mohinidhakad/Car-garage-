package Zopsmart_pro

import (
	"api/handler"
	"api/store"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := store.New()
	h := handler.New(s)

	app.Server.ValidateHeaders = false

	// specifying the different routes supported by this service
	app.GET("/car", h.Get)
	app.POST("/car", h.Create)

	app.Server.HTTP.Port = 3306
	app.Server.MetricsPort = 2113

	app.Start()
}
