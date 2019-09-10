## echoの使い方メモ  

## ミドルウェア  
```go:
import "github.com/labstack/echo/middleware"

e := echo.New()

// ミドルウェアの追加
e.Use(middleware.Logger())
```

### middleweare.Logger()  
画面にログを表示

## 環境構築  
`go get -u github.com/labstack/echo/...`

## 実行方法  
`go run server.go`  
http://localhost:1323
