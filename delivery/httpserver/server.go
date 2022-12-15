package httpserver

import (
	"log"

	"github.com/Dorogobid/accounting-service/internal/domain/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	e *echo.Echo
	useCases *usecases.UseCases
}

func New(uc *usecases.UseCases) *HttpServer {
	return &HttpServer{e: echo.New(), useCases: uc}
}

// Start starting http server
func (h *HttpServer) Start() {
	h.registerHandlers()
	h.configureLogger()
	// h.registerTemplates(e)

	log.Fatal(h.e.Start(":8080"))
}

func (h *HttpServer) registerHandlers() {
	h.e.GET("/", h.indexHandler)
	h.e.GET("/clients", h.clientsHandler)
}

func (h *HttpServer) configureLogger() {
	h.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	  }))
}