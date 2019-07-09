package router

import (
  "../handler"
  "github.com/labstack/echo"
)

func Router(e *echo.Echo) *echo.Echo {

  e.GET("/hello", handler.HelloHandler)
  e.GET("/:id", handler.IdHandler) // urlの値を読み取り表示する
  e.GET("/plus/:id", handler.IdPlusHandler) 
  return e
}
