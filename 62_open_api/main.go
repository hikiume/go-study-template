package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct{}

func (s *Server) GetPing(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func main() {
	e := echo.New()

	server := &Server{}

	RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8080"))
}
