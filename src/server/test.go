package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Usuario struct {
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

var usuarios = []Usuario{}
var validate *validator.Validate

func main() {
	r := gin.Default()
	validate = validator.New() // Inicializa o validador

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Funcionou! Agora tente acessar uma rota qualquer, tipo 'localhost:8080/João'")
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "------> %s <------", name)
	})

	r.POST("/register", register)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func register(c *gin.Context) {
	var novoUsuario Usuario

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&novoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valida o JSON recebido de acordo com as regras definidas na struct Usuario
	if err := validate.Struct(novoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Checa se o usuário já existe
	for _, usuario := range usuarios {
		if usuario.Nome == novoUsuario.Nome || usuario.Email == novoUsuario.Email {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário com este nome ou email já existe"})
			return
		}
	}

	usuarios = append(usuarios, novoUsuario)
	c.IndentedJSON(http.StatusCreated, novoUsuario)
}
