package server

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/zombeer/go-uniswap-info/uniswap"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, WEB3 World!")
	})
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("server/templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/api/pair/:address", handleApiGetPair)
	e.GET("/pair/:address", handleGetPair)

	e.Logger.Fatal(e.Start(":1234"))
}

func handleApiGetPair(c echo.Context) error {
	address := c.Param("address")
	pairInfo := uniswap.GetPairInfo(address)
	candles := uniswap.PricesToCandles(pairInfo.Prices)
	return c.JSONPretty(http.StatusOK, &candles, "  ")
}

func handleGetPair(c echo.Context) error {
	address := c.Param("address")
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"pairAddress": address,
	})
}
