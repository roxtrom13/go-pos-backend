package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/roier/POS/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creating a new conection to database
func SetDatabase() *gorm.DB {
  envErr := godotenv.Load()
  if envErr != nil {
    panic("Failed to load env file")
  }

  user := os.Getenv("DB_USER")
  pass := os.Getenv("DB_PASS")
  host := os.Getenv("DB_HOST")
  name := os.Getenv("DB_NAME")

  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
    host,
    user,
    pass,
    name,
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Failed to connect database")
  }
  db.AutoMigrate(&models.Product{}, &models.User{})
  return db
}

func CloseDatabase(db *gorm.DB) {
  dbPsql, err := db.DB()
  if err != nil {
    panic("Failed to close database connection")
  }

  dbPsql.Close()
}
