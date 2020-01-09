package model

type Order struct {
	OderNumber string `gorm:"column:orderNumber"`
	UserID     int64  `gorm:"column:userID"`
}
