package model

type User struct {
	UserID   int64  `gorm:"column:userID;primary_key"`
	UserName string `gorm:"column:userName"`
	TrueName string `gorm:"column:trueName"`
	NickName string `gorm:"column:nickName"`
}
