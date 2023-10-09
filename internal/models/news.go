package models

type News struct {
	Id         int    `gorm:"primaryKey"`
	Title      string `gorm:"not null" validate:"required_without_all=Content Categories,min=5"`
	Content    string `gorm:"not null" validate:"required_without_all=Title Categories,min=5"`
	Categories []int  `gorm:"-" validate:"required_without_all=Title Content"`
}
