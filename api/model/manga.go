package model

import "github.com/gin-gonic/gin"

type Manga struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Year   uint   `json:"year" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"primaryKey"`
	Author string `json:"author" gorm:"primaryKey"`
}

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, response)
}
