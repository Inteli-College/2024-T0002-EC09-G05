package platform

import (
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type GetInput struct {
	Id int `json:"id"`
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
	pg.Raw("SELECT elements.name, layouts.index, props.value FROM users JOIN layouts ON users.id= layouts.user_refer  JOIN elements ON layouts.elements_refer = elements.ID JOIN props ON elements.props_refer = props.ID WHERE users.ID = ? ORDER BY layouts.index ASC", input.Id).Scan(&result)

	c.JSON(200, gin.H{"status": "ok", "elements": result})

}
