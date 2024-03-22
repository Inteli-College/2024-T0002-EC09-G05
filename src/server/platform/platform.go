package platform

import (
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	// needed to validate if the incoming data is correct
)

type GetInput struct {
	Id    int    `json:"id"`
	Token string `json:"Token" validate:"required"`
}

func GetAll(c *gin.Context, pg *gorm.DB) {

	var validate = validator.New()
	var input GetInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(input)
	if validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	type Result struct {
		Name  string
		Index int
		Value string
	}

	var result []Result
	fmt.Fprintln(os.Stdout, []any{input.Id}...)
	// Raw SQL
	pg.Raw("SELECT elements.name, props.value, elements.index FROM users INNER JOIN elements ON users.id= elements.user_refer INNER JOIN props ON elements.props_refer= props.id WHERE users.id = ?", input.Id).Scan(&result)

	c.JSON(200, gin.H{"status": "ok", "elements": result, "id": input.Id})

}
