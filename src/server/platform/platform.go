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

func GetAllComponents(c *gin.Context, pg *gorm.DB) {

	type Users_raw struct {
		ID          uint
		Name        string
		Email       string
		Role        uint
		Directorate string
	}

	type Id struct {
		Id uint
	}

	type Props struct {
		Value string
	}

	type Elements struct {
		Name string
	}

	type Directorates struct {
		Directorate string
	}

	var id []Id
	var props []Props
	var elements []Elements
	var directorates []Directorates
	var users_raw []Users_raw

	// Raw SQL
	pg.Raw("SELECT users.id, users.name, users.email, users.role, directorates.directorate    FROM users JOIN directorates ON users.directorate_refer = directorates.id ORDER BY users.id ASC").Scan(&users_raw)
	pg.Raw("SELECT elements.name FROM elements").Scan(&elements)
	pg.Raw("SELECT directorates.directorate FROM directorates").Scan(&directorates)
	pg.Raw("SELECT props.value FROM props").Scan(&props)
	pg.Raw("SELECT users.id FROM users").Scan(&id)

	c.JSON(200, gin.H{
		"status":       "ok",
		"users_id":     id,
		"props":        props,
		"elements":     elements,
		"directorates": directorates,
		"users_raw":    users_raw,
	})

}
