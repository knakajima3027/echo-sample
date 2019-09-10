package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "strconv"
)


func HelloHandler(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}


// urlの値を読み取り表示する
func IdHandler(c echo.Context) error {
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}


// urlの値を読み取りidの値に1を加えた数を表示する
func IdPlusHandler(c echo.Context) error {
    id := c.Param("id")

    int_id, _ := strconv.Atoi(id)
    int_id += 1

    return c.String(http.StatusOK, strconv.Itoa(int_id))
}


// jsonファイルを返す
func JsonHandler (c echo.Context) error {
  jsonMap := map[string]string{
            "foo": "bar",
            "hoge": "fuga",
        }

  return c.JSON(http.StatusOK, jsonMap)
}
