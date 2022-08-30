package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zombeer/go-uniswap-info/uniswap"
)

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, WEB3 World!")
	})
	e.GET("/pair/:address", handleGetPair)
	e.Logger.Fatal(e.Start(":1234"))
}

func handleGetPair(c echo.Context) error {
	address := c.Param("address")
	pairInfo := uniswap.GetPairInfo(address)
	return c.JSONPretty(http.StatusOK, &pairInfo, "  ")
}
