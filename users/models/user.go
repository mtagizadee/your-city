package models

type User struct {
  Id int `gorm:"primaryKey;autoIncrement"`
  Name string `gorm:"not null"`
  Surname string `gorm:"not null"`
  Email string `gorm:"unique;not null"`
  Password string `gorm:"not null"`
}
