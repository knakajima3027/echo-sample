## echoの使い方メモ  

### ミドルウェア  
```go:
import "github.com/labstack/echo/middleware"

e := echo.New()

// ミドルウェアの追加
e.Use(middleware.Logger())
```

### middleweare.Logger()  
画面にログを表示

### echoのインストール  
`go get -u github.com/labstack/echo/...`

### 実行方法  
`go run server.go`  
http://localhost:1323


## gormの使い方メモ  
### gormのインストール  
`go get github.com/jinzhu/gorm`  

### SQLite3ドライバのインストール  
`go get github.com/mattn/go-sqlite3`

## コンテナの生成  
`docker-compose up`
