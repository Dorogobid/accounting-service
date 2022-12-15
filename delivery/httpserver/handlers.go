package httpserver

import (
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HttpServer) indexHandler(c echo.Context) error {
	h.setTemplate("./templates/layout.html", "./templates/index.html")

	return c.Render(http.StatusOK, "", nil)
}

func (h *HttpServer) clientsHandler(c echo.Context) error {
	h.setTemplate("./templates/layout.html", "./templates/clients.html")

	clients, err := h.useCases.GetClientsWithType()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.Render(http.StatusOK, "", clients)
}

func (h *HttpServer) setTemplate(lFileName string, tFilename string) {
	tmpl := template.Must(template.ParseFiles(lFileName, tFilename))
	t := &Template{
		templates: tmpl,
	}
	h.e.Renderer = t
}