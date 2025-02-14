package utils

import (
	"go_mangasnake_api/api/model"

	"github.com/gin-gonic/gin"
)

func RespondWithError(ctx *gin.Context, status int, message string, data any) {
	model.ResponseJSON(ctx, status, message, data)
}
