package config

import "fmt"

type database struct {
  Host string
  Port string
  User string
  Password string
  Name string

  Dsn string // complete uri to connect to the database
}

var databaseConfig database

func loadDatabase(host, port, user, password, name string) {
  databaseConfig.Host = host
  databaseConfig.Port = port
  databaseConfig.User = user
  databaseConfig.Password = password
  databaseConfig.Name = name

  databaseConfig.Dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
}

func GetDatabaseConfig() database {
  return databaseConfig
}
