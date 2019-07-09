package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./router"
)




func main() {
  e := echo.New()

  router.Router(e)

  // ミドルウェアの追加
  e.Use(middleware.Logger())


  e.Start(":1323")
}
