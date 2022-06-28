package models

type User struct {
	Id       uint   `gorm:"primaryKey;auto_increment;not_null"`
	Name     string `gorm:"not_null"`
	Email    string
	Password []byte `gorm:"not_null"`

}
