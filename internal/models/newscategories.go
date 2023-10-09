package models

type NewsCategories struct {
	NewsId     int `gorm:"primaryKey"`
	CategoryId int `gorm:"primaryKey"`
}
