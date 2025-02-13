package model

type List struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	UserID uint    `json:"user_id"`
	User   User    `json:"user" gorm:"foreignKey:UserID"`
	Mangas []Manga `json:"mangas" gorm:"many2many:list_mangas"`
}
