package models

type Product struct {
	Pid     uint   `json:"pid" gorm:"primaryKey;auto_increment;not_null"`
	Pname   string  `json:"pname"`
	Descript string `json:"descript"`
	Price   string `json:"price"`
	Picture string  `json:"picture"`
}
