package model

type Manga struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Year        uint     `json:"year" gorm:"primaryKey"`
	Title       string   `json:"title" gorm:"primaryKey"`
	Author      string   `json:"author" gorm:"primaryKey"`
	Genres      []string `json:"genres" gorm:"serializer:json"`
	Status      string   `json:"status" gorm:"primaryKey"`
	Type        string   `json:"type" gorm:"primaryKey"`
	Description string   `json:"description" gorm:"primaryKey"`
}
