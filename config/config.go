package config

import (
  "os"
  "github.com/joho/godotenv"
)

func LoadConfig() {
  godotenv.Load()

  // load server config
  loadServer(os.Getenv("SERVER_PORT"), os.Getenv("SERVER_HOST"))
}
