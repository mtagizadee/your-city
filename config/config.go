package config

import (
  "os"
  "github.com/joho/godotenv"
)

func LoadConfig() {
  godotenv.Load()

  // load server config
  loadServer(os.Getenv("SERVER_PORT"), os.Getenv("SERVER_HOST"))

  // load database config
  loadDatabase(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
