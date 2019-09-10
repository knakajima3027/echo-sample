package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./handler"
)


func main() {

  // Echoインスタンスの生成
  e := echo.New()

  // ミドルウェアの追加
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルーティング
  e.GET("/hello",    handler.HelloHandler)
  e.GET("/userid/:id",      handler.IdHandler) // urlの値を読み取り表示する
  e.GET("/userid/plus/:id", handler.IdPlusHandler)
  e.GET("/json",     handler.JsonHandler)

  e.Logger.Fatal(e.Start(":1323"))
}
