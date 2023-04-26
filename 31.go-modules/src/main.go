package main

import "net/http"

//template
func main() {

	e := echo.new()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Echo!")
	})
	e.logger.Fatal(e.start(":1323"))

}
