package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./router"
)


func main() {
  e := echo.New()

  // ミドルウェアの追加
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  router.Router(e)

  e.Start(":1323")
}
