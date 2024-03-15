package auth

import (
	"g5/server/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	// needed to validate if the incoming data is correct
)

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context, pg *gorm.DB) {

	var validate = validator.New()
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(input)
	if validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	token, err := LoginCheck(input.Email, input.Password, pg)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{"token": token})

}

func Register(c *gin.Context, pg *gorm.DB) {

	var validate = validator.New()
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(input)
	if validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	tx := pg.First(&db.User{}, "email = ?", input.Email)

	if tx.RowsAffected > 0 {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword := HashPassword(input.Password)

	u := db.User{
		Email:    input.Email,
		Password: hashedPassword,
		Name:     input.Name,
	}

	
	u.SaveUser(pg, &u)
	
	token, _ := GenerateToken(u.ID, u.Role)

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"status": "ok", "token": token})

}

func ChangeRole(c *gin.Context, pg *gorm.DB) {

	requesterRole := c.GetInt("role")

	if requesterRole != 2 {
		c.JSON(400, gin.H{"error": "You don't have the necessary role to access this resource"})
		return
	}
	
	var validate = validator.New()
	var input struct {
		Email string `json:"email" validate:"required,email"`
		Role  string `json:"role" validate:"required, eq=1|eq=2"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(input)
	if validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	tx := pg.First(&db.User{}, "email = ?", input.Email)

	if tx.RowsAffected == 0 {
		c.JSON(400, gin.H{"error": "Email not found"})
		return
	}

	pg.Model(&db.User{}).Where("email = ?", input.Email).Update("role", input.Role)

	c.JSON(200, gin.H{"status": "ok"})

}