package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func init() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    s := standard.New("")
    s.SetHandler(e)
    http.Handle("/", s)
}
