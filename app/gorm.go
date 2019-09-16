package main

import (
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

// テーブルのカラムを定義
type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {

  // データベースへの接続
  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("データベースへの接続に失敗しました")
  }
  defer db.Close()

  // スキーマのマイグレーション
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "L1212", Price: 1000})
  db.Create(&Product{Code: "R2525", Price: 1192})
  db.Create(&Product{Code: "R25251", Price: 1192})

  var product Product
  var products []Product // 複数のデータを格納したいときに利用

  // Read
  db.First(&product, 1) // idが1の製品を探す
  db.First(&product, "code = ?", "L1212") // codeがL1212の製品を探す
  db.Find(&products)  // 全データの取得
  fmt.Println(products)

  // Update - 製品価格を2,000に更新します
  db.Model(&product).Update("Price", 2000)

  // Delete - 製品を削除します
  db.Delete(&product)

}
