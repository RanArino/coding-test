package main

import (
  "fmt"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // find product with integer primary key
  fmt.Printf("Found product: %+v\n", product)
  db.First(&product, "code = ?", "D42") // find product with code D42
  fmt.Printf("Found product: %+v\n", product)

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  fmt.Printf("Updated product: %+v\n", product)

  // Delete - delete product
  db.Delete(&product, 1)
  fmt.Println("Deleted product")
}
