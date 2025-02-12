package handlers

import (
	"go_mangasnake_api/api/middleware"
	"go_mangasnake_api/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	// Verificar se o e-mail já está cadastrado
	var existingUser model.User
	if err := DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		model.ResponseJSON(c, http.StatusConflict, "Email already registered", nil)
		return
	}
	if err := user.HashPassword(); err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		model.ResponseJSON(c, http.StatusInternalServerError, "Could not hash password", nil)
		return
	}
	if err := DB.Create(&user).Error; err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		model.ResponseJSON(c, http.StatusBadRequest, "Could not create user", nil)
		return
	}

	//DB.Create(&user)
	model.ResponseJSON(c, http.StatusCreated, "User created successfully!!!", user)
}

func LoginUser(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	var user model.User
	if err := DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	if !user.CheckPassword(request.Password) {
		model.ResponseJSON(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	token, err := middleware.GenerateJWTU(user.ID)
	if err != nil {
		model.ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}

	model.ResponseJSON(c, http.StatusOK, "Login successful", gin.H{"token": token})
}
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Verifica se o usuário existe
	var user model.User
	if err := DB.First(&user, userID).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "User not found", nil)
		return
	}

	// Remove o usuário
	if err := DB.Delete(&user).Error; err != nil {
		model.ResponseJSON(c, http.StatusInternalServerError, "Could not delete user", nil)
		return
	}

	model.ResponseJSON(c, http.StatusOK, "User deleted successfully", nil)

	// jeito antigo seguindo uma linha de raciocinio que TALVEZ fique mais bonito o codigo, porém tava dando erro de duplicata, guardar pra mais pra frente ver se consigo arrumar
	// userID, exists := c.Get("user_id")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 	return
	// }

	// if err := DB.Delete(&model.User{}, userID).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
func GetUser() {

}

func GetUsers() {

}
